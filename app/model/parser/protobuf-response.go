package parser

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/emicklei/proto"
	"github.com/kyleu/npn/app/model/schema"
	"github.com/kyleu/npn/app/util"
	"path"
	"path/filepath"
	"strings"
	"text/scanner"
)

type ProtobufResponse struct {
	RootFile string         `json:"root"`
	Data     []interface{}  `json:"data"`
	Schema   *schema.Schema `json:"schema"`
}

func NewProtobufResponse(key string) *ProtobufResponse {
	md := schema.Metadata{Comments: nil, Origin: schema.OriginProtobuf, Source: util.FilenameOf(key)}
	return &ProtobufResponse{
		RootFile: key,
		Data:     make([]interface{}, 0),
		Schema:   schema.NewSchema(key, []string{key}, &md),
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
	model := &schema.Model{Key: msg.Name, Fields: nil, Metadata: md}
	for _, el := range msg.Elements {
		switch f := el.(type) {
		case *proto.NormalField:
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			t := getProtoType(f.Type, f.Repeated, f.Options)
			err := model.AddField(&schema.Field{Key: f.Name, Type: t, Optional: f.Optional, Metadata: md})
			if err != nil {
				return errors.Wrap(err, "can't add field")
			}
		case *proto.Oneof:
			oo, err := s.addOneOf(f)
			if err != nil {
				return errors.Wrap(err, "can't register union")
			}
			err = model.AddField(&schema.Field{Key: f.Name, Type: oo.Key, Metadata: md})
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
			t := getProtoType(f.Type, false, f.Options)
			v = append(v, &schema.Field{Key: f.Name, Type: t, Optional: false, Metadata: md})
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
	methods := make([]*schema.Method, 0)
	for _, el := range svc.Elements {
		switch f := el.(type) {
		case *proto.RPC:
			md := getProtobufMetadata(f.Position, f.Comment, f.InlineComment)
			args := schema.Fields{{Key: "arg", Type: f.RequestType}}
			methods = append(methods, &schema.Method{Key: f.Name, Args: args, Ret: f.ReturnsType, Metadata: md})
		default:
			return errors.New(fmt.Sprintf("can't add service field of type [%T]", el))
		}
	}

	return s.Schema.AddService(&schema.Service{Key: svc.Name, Methods: methods, Metadata: md})
}

func getProtobufMetadata(pos scanner.Position, comments ...*proto.Comment) *schema.Metadata {
	cmt := make([]string, 0)
	for _, cs := range comments {
		if cs != nil {
			for _, l := range cs.Lines {
				l = strings.TrimSpace(l)
				cmt = append(cmt, l)
			}
		}
	}
	return &schema.Metadata{
		Comments: cmt,
		Origin:   schema.OriginProtobuf,
		Source:   util.FilenameOf(pos.Filename),
		Line:     pos.Line,
		Column:   pos.Column - 1,
	}
}

func getProtoType(t string, repeated bool, options []*proto.Option) string {
	for _, opt := range options {
		println(fmt.Sprintf("option [%v] provided for proto type", opt))
	}
	if repeated {
		return "List[" + t + "]"
	}
	return t
}
