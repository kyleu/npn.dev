package parseprotobuf

import (
	"emperror.dev/errors"
	"github.com/emicklei/proto"
	"os"
	"path"
)

func (p *ProtobufParser) ParseProtobufFile(paths []string) (*ProtobufResponse, error) {
	return p.parse(paths, NewProtobufResponse(paths))
}

func (p *ProtobufParser) parse(paths []string, ret *ProtobufResponse) (*ProtobufResponse, error) {
	rsp := ret
	var err error
	for _, pth := range paths {
		rsp, err = p.parsePath(pth, rsp)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing protobuf")
		}
	}
	return rsp, nil
}

func (p *ProtobufParser) parsePath(fn string, ret *ProtobufResponse) (*ProtobufResponse, error) {
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
				_, err = p.parsePath(path.Join(path.Dir(fn), imp.Filename), ret)
				p.log(ret, err)
			}
		}),
		proto.WithOption(func (opt *proto.Option) { p.log(ret, ret.addOption(opt)) }),
		proto.WithEnum(func (en *proto.Enum) { p.log(ret, ret.addEnum(en)) }),
		proto.WithMessage(func (msg *proto.Message) { p.log(ret, ret.addMessage(msg)) }),
		proto.WithService(func (svc *proto.Service) { p.log(ret, ret.addService(svc)) }),
	)
	ret.currPkg = nil
	return ret, nil
}
