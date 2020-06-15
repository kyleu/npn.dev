package parser

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/emicklei/proto"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/model/schema/schematypes"
	"github.com/kyleu/npn/app/util"
	"path"
	"path/filepath"
)

type ProtobufResponse struct {
	RootFile string         `json:"root"`
	Data     []interface{}  `json:"data"`
	Schema   *schema.Schema `json:"schema"`
}

func NewProtobufResponse(paths []string) *ProtobufResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginProtobuf, Source: paths[0]}
	return &ProtobufResponse{
		RootFile: paths[0],
		Data:     make([]interface{}, 0),
		Schema:   schema.NewSchema(paths[0], paths, &md),
	}
}

func (s *ProtobufResponse) addPackage(pkg *proto.Package) error {
	md := getProtobufMetadata(pkg.Position, pkg.Comment, pkg.InlineComment)
	ret, err := util.StringKeyMapFromPairs("t", "package", "name", pkg.Name, "metadata", md)
	s.Data = append(s.Data, ret)
	return err
}

func (s *ProtobufResponse) addImport(imp *proto.Import) (bool, error) {
	md := getProtobufMetadata(imp.Position, imp.Comment, imp.InlineComment)
	ret, _ := util.StringKeyMapFromPairs("t", "import", "kind", imp.Kind, "filename", imp.Filename, "metadata", md)
	s.Data = append(s.Data, ret)
	p := path.Join(filepath.Dir(s.RootFile), imp.Filename)
	return s.Schema.AddPath(p), nil
}

func (s *ProtobufResponse) addOption(opt *proto.Option) error {
	md := getProtobufMetadata(opt.Position, opt.Comment, opt.InlineComment)
	ret, err := util.StringKeyMapFromPairs("t", "option", "name", opt.Name, "value", opt.Constant.Source, "metadata", md)
	s.Data = append(s.Data, ret)
	return err
}

func (s *ProtobufResponse) addEnum(en *proto.Enum) error {
	md := getProtobufMetadata(en.Position, en.Comment)
	vals := make(schema.EnumValues, 0)
	for _, v := range en.Elements {
		switch ef := v.(type) {
		case *proto.EnumField:
			efmd := getProtobufMetadata(ef.Position, ef.Comment, ef.InlineComment)
			vals = append(vals, &schema.EnumValue{Key: ef.Name, IntVal: ef.Integer, Metadata: efmd})
		default:
			return errors.New(fmt.Sprintf("can't add enum field of type [%T]", v))
		}
	}
	return s.Schema.AddEnum(&schema.Enum{Key: en.Name, Values: vals, Metadata: md})
}

func (s *ProtobufResponse) addMessage(msg *proto.Message) error {
	md := getProtobufMetadata(msg.Position, msg.Comment)
	model := &schema.Model{Key: msg.Name, Type: schema.ModelTypeStruct, Fields: nil, Metadata: md}
	for _, el := range msg.Elements {
		switch f := el.(type) {
		case *proto.NormalField:
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			typ := getProtobufType(f.Type, f.Optional, f.Repeated, f.Options)
			err := model.AddField(&schema.Field{Key: f.Name, Type: typ, Metadata: md})
			if err != nil {
				return errors.Wrap(err, "can't add field")
			}
		case *proto.Oneof:
			oo, err := s.addOneOf(f)
			if err != nil {
				return errors.Wrap(err, "can't register union")
			}
			typ := schematypes.Wrap(schematypes.Unknown{X: oo.Key})
			err = model.AddField(&schema.Field{Key: f.Name, Type: typ, Metadata: md})
			if err != nil {
				return errors.Wrap(err, "can't add field")
			}
		default:
			return errors.New(fmt.Sprintf("can't add message field of type [%T]", el))
		}
	}
	return s.Schema.AddModel(model)
}

func (s *ProtobufResponse) addOneOf(oo *proto.Oneof) (*schema.Union, error) {
	v := make([]*schema.Field, 0, len(oo.Elements))
	for _, el := range oo.Elements {
		switch f := el.(type) {
		case *proto.OneOfField:
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			typ := getProtobufType(f.Type, false, false, f.Options)
			v = append(v, &schema.Field{Key: f.Name, Type: typ, Metadata: md})
		default:
			return nil, errors.New(fmt.Sprintf("can't add union field of type [%T]", el))
		}
	}
	md := getProtobufMetadata(oo.Position, oo.Comment)
	u := &schema.Union{Key: oo.Name, Variants: v, Metadata: md}
	err := s.Schema.AddUnion(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *ProtobufResponse) addService(svc *proto.Service) error {
	md := getProtobufMetadata(svc.Position, svc.Comment)
	methods := make([]*schema.Field, 0)
	for _, el := range svc.Elements {
		switch f := el.(type) {
		case *proto.RPC:
			typ := schematypes.Method{
				Args: schematypes.Arguments{{Key: "arg", Type: getTypeForProtobufName(f.RequestType)}},
				Ret:  getTypeForProtobufName(f.ReturnsType),
			}
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			methods = append(methods, &schema.Field{Key: f.Name, Type: schematypes.Wrap(typ), Metadata: md})
		default:
			return errors.New(fmt.Sprintf("can't add service field of type [%T]", el))
		}
	}

	return s.Schema.AddModel(&schema.Model{Key: svc.Name, Type: schema.ModelTypeService, Fields: methods, Metadata: md})
}
