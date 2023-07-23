package ast

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"members/storage"
	"os"

	"github.com/emicklei/proto"
)

type (
	readerFactory = func() (io.Reader, error)

	event struct {
		enum    *proto.Enum
		imports *proto.Import
		message *proto.Message
		normal  *proto.NormalField
		oneof   *proto.Oneof
		option  *proto.Option
		svc     *proto.Service
		pkg     *proto.Package
		rpc     *proto.RPC
	}

	protoView struct {
		get_reader readerFactory
		seen       map[string]event
		events     chan *event
	}

	comparableProto struct {
		prev *protoView
		this *protoView
	}
)

func (pr *protoView) Upload(store storage.Store) error {
	return nil
}

func (pr *protoView) Parse() (*proto.Proto, error) {
	reader, err := pr.get_reader()
	if err != nil {
		return nil, err
	}
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return definition, nil
}

func (pr *protoView) Watch() {
	for ev := range pr.events {
		log.Printf("%+v", ev)
		switch {
		case ev.enum != nil:
			pr.seen[ev.enum.Name] = *ev
		}
	}
}

func (pr *protoView) Walk() error {
	defn, err := pr.Parse()
	if err != nil {
		return err
	}
	go pr.Watch()
	proto.Walk(defn,
		proto.WithEnum(pr.handleEnum),
		proto.WithImport(pr.handleImports),
		proto.WithMessage(pr.handleMessage),
		proto.WithNormalField(pr.handleNormalField),
		proto.WithOneof(pr.handleOneOf),
		proto.WithOption(pr.handleOption),
		proto.WithService(pr.handleService),
		proto.WithPackage(pr.handlePackage),
		proto.WithRPC(pr.handleRpc),
	)
	return nil
}

func FromFile(pth string) *protoView {
	bts, err := os.ReadFile(pth)
	return NewViewer(func() (io.Reader, error) {
		if err != nil {
			return nil, fmt.Errorf("failed to open %s", pth)
		}
		return bytes.NewReader(bts), nil
	})
}

func FromBinary(bts []byte) *protoView {
	return NewViewer(func() (io.Reader, error) {
		reader := bytes.NewReader(bts)
		return reader, nil
	})
}

func NewViewer(get_reader readerFactory) *protoView {
	return &protoView{
		get_reader,
		make(map[string]event),
		make(chan *event),
	}
}

func (pr *protoView) handleService(s *proto.Service) {
	pr.events <- &event{
		svc: s,
	}
}
func (pr *protoView) handleRpc(s *proto.RPC) {
	pr.events <- &event{
		rpc: s,
	}
}
func (pr *protoView) handleEnum(s *proto.Enum) {
	pr.events <- &event{
		enum: s,
	}
}
func (pr *protoView) handleNormalField(s *proto.NormalField) {
	pr.events <- &event{
		normal: s,
	}
}
func (pr *protoView) handleImports(s *proto.Import) {
	log.Printf("import : %+v", s)
	pr.events <- &event{
		imports: s,
	}
}
func (pr *protoView) handleOneOf(s *proto.Oneof) {
	pr.events <- &event{
		oneof: s,
	}
}
func (pr *protoView) handlePackage(s *proto.Package) {
	pr.events <- &event{
		pkg: s,
	}
}

func (pr *protoView) handleOption(s *proto.Option) {
	pr.events <- &event{
		option: s,
	}
}
func (pr *protoView) handleMessage(m *proto.Message) {
	pr.events <- &event{
		message: m,
	}

	lister := new(optionLister)
	for _, each := range m.Elements {
		each.Accept(lister)
	}

	for n, ele := range m.Elements {
		log.Printf("msg ele %d: %+v", n, ele)
	}

}

type optionLister struct {
	proto.NoopVisitor
}

func (l optionLister) VisitOption(o *proto.Option) {
	log.Printf("opt: %+v", *o)
}
