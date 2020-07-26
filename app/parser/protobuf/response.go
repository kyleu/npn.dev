package parseprotobuf

import (
	"fmt"
	"github.com/kyleu/npn/npncore"
	"path"
	"path/filepath"
	"strings"

	parseutil "github.com/kyleu/npn/app/parser/util"

	"emperror.dev/errors"
	"github.com/emicklei/proto"
	"github.com/kyleu/npn/app/schema"
	"github.com/kyleu/npn/app/schema/schematypes"
	"github.com/kyleu/npn/app/util"
)

type ProtobufResponse struct {
	Rsp      *parseutil.ParseResponse
	rootFile string
	currPkg  util.Pkg
}

func NewProtobufResponse(paths []string) *ProtobufResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginProtobuf, Source: paths[0]}
	return &ProtobufResponse{Rsp: parseutil.NewParseResponse(paths, md)}
}

func (s *ProtobufResponse) addPackage(pkg *proto.Package) error {
	md := getProtobufMetadata(pkg.Position, pkg.Comment, pkg.InlineComment)
	ret, err := npncore.StringKeyMapFromPairs("t", "package", "name", pkg.Name, "metadata", md)
	s.Rsp.Data = append(s.Rsp.Data, ret)
	s.currPkg = strings.Split(pkg.Name, ".")
	return err
}

func (s *ProtobufResponse) addImport(imp *proto.Import) bool {
	md := getProtobufMetadata(imp.Position, imp.Comment, imp.InlineComment)
	ret, _ := npncore.StringKeyMapFromPairs("t", "import", "kind", imp.Kind, "filename", imp.Filename, "metadata", md)
	s.Rsp.Data = append(s.Rsp.Data, ret)
	p := path.Join(filepath.Dir(s.rootFile), imp.Filename)
	return s.Rsp.Schema.AddPath(p)
}

func (s *ProtobufResponse) addOption(opt *proto.Option) error {
	md := getProtobufMetadata(opt.Position, opt.Comment, opt.InlineComment)
	ret, err := npncore.StringKeyMapFromPairs("t", "option", "name", opt.Name, "value", opt.Constant.Source, "metadata", md)
	s.Rsp.Data = append(s.Rsp.Data, ret)
	return err
}

func (s *ProtobufResponse) addEnum(en *proto.Enum) error {
	md := getProtobufMetadata(en.Position, en.Comment)
	vals := make(schema.Fields, 0, len(en.Elements))
	for _, v := range en.Elements {
		switch ef := v.(type) {
		case *proto.EnumField:
			efmd := getProtobufMetadata(ef.Position, ef.Comment, ef.InlineComment)
			vals = append(vals, &schema.Field{Key: ef.Name, Type: schematypes.Wrap(schematypes.EnumValue{}), Metadata: efmd})
		default:
			return errors.New(fmt.Sprintf("can't add enum field of type [%T]", v))
		}
	}
	return s.Rsp.Schema.AddModel(&schema.Model{Key: en.Name, Pkg: s.currPkg, Type: schema.ModelTypeEnum, Fields: vals, Metadata: md})
}

func (s *ProtobufResponse) addMessage(msg *proto.Message) error {
	md := getProtobufMetadata(msg.Position, msg.Comment)
	model := &schema.Model{Key: msg.Name, Pkg: s.currPkg, Type: schema.ModelTypeStruct, Fields: nil, Metadata: md}
	for _, el := range msg.Elements {
		switch f := el.(type) {
		case *proto.NormalField:
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			typ := getProtobufType(s.currPkg, f.Type, f.Optional, f.Repeated, f.Options)
			err := model.AddField(&schema.Field{Key: f.Name, Type: typ, Metadata: md})
			if err != nil {
				return errors.Wrap(err, "can't add field")
			}
		case *proto.Oneof:
			oo, err := s.addOneOf(f)
			if err != nil {
				return errors.Wrap(err, "can't register union")
			}
			typ := schematypes.ReferenceWrapped(s.currPkg, oo.Key)
			err = model.AddField(&schema.Field{Key: f.Name, Type: typ, Metadata: md})
			if err != nil {
				return errors.Wrap(err, "can't add field")
			}
		default:
			return errors.New(fmt.Sprintf("can't add message field of type [%T]", el))
		}
	}
	return s.Rsp.Schema.AddModel(model)
}

func (s *ProtobufResponse) addOneOf(oo *proto.Oneof) (*schema.Model, error) {
	v := make(schema.Fields, 0, len(oo.Elements))
	for _, el := range oo.Elements {
		switch f := el.(type) {
		case *proto.OneOfField:
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			typ := getProtobufType(s.currPkg, f.Type, false, false, f.Options)
			v = append(v, &schema.Field{Key: f.Name, Type: typ, Metadata: md})
		default:
			return nil, errors.New(fmt.Sprintf("can't add union field of type [%T]", el))
		}
	}
	md := getProtobufMetadata(oo.Position, oo.Comment)
	u := &schema.Model{Key: oo.Name, Pkg: s.currPkg, Type: schema.ModelTypeUnion, Fields: v, Metadata: md}
	err := s.Rsp.Schema.AddModel(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *ProtobufResponse) addService(svc *proto.Service) error {
	md := getProtobufMetadata(svc.Position, svc.Comment)
	methods := make(schema.Fields, 0)
	for _, el := range svc.Elements {
		switch f := el.(type) {
		case *proto.RPC:
			typ := schematypes.Method{
				Args: schematypes.Arguments{{Key: "arg", Type: getTypeForProtobufName(s.currPkg, f.RequestType)}},
				Ret:  getTypeForProtobufName(s.currPkg, f.ReturnsType),
			}
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			methods = append(methods, &schema.Field{Key: f.Name, Type: schematypes.Wrap(typ), Metadata: md})
		default:
			return errors.New(fmt.Sprintf("can't add service field of type [%T]", el))
		}
	}

	return s.Rsp.Schema.AddModel(&schema.Model{Key: svc.Name, Pkg: s.currPkg, Type: schema.ModelTypeService, Fields: methods, Metadata: md})
}
