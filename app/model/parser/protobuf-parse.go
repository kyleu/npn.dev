package parser

import (
	"emperror.dev/errors"
	"github.com/emicklei/proto"
	"os"
	"path"
)

func (p *ProtobufParser) ParseProtobufFile(path string) (*ProtobufResponse, error) {
	return p.parse(path, NewProtobufResponse(path))
}

func (p *ProtobufParser) parse(fn string, ret *ProtobufResponse) (*ProtobufResponse, error) {
	reader, err := os.Open(fn)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file [" +fn+ "]")
	}
	defer func() { _ = reader.Close() }()

	parser := proto.NewParser(reader)
	parser.Filename(fn)
	definition, err := parser.Parse()
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse protobuf file at [" +fn+ "]")
	}

	proto.Walk(
		definition,
		proto.WithPackage(func (pkg *proto.Package) { p.log(ret, ret.addPackage(pkg))}),
		proto.WithImport(func (imp *proto.Import) {
			newFile, err := ret.addImport(imp)
			p.log(ret, err)
			if newFile {
				_, err = p.parse(path.Join(path.Dir(fn), imp.Filename), ret)
				p.log(ret, err)
			}
		}),
		proto.WithOption(func (opt *proto.Option) { p.log(ret, ret.addOption(opt)) }),
		proto.WithEnum(func (en *proto.Enum) { p.log(ret, ret.addEnum(en)) }),
		proto.WithMessage(func (msg *proto.Message) { p.log(ret, ret.addMessage(msg)) }),
		// proto.WithOneof(func (oo *proto.Oneof) { p.log(ret, ret.addOneOf(oo)) }),
		// proto.WithRPC(func (rpc *proto.RPC) { p.log(ret, ret.addRPC(rpc)) }),
		proto.WithService(func (svc *proto.Service) { p.log(ret, ret.addService(svc)) }),
	)

	return ret, nil
}
