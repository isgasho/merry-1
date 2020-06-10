package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/igso/languages/manifest"
	"strings"
)

type KeyValue struct {
	Key   string
	Value string
}

type JavaPackage struct {
	Name         string
	VersionInfo  string
	StartVersion string
	EndVersion   string
	Config       []KeyValue
}

type IgsoManifest struct {
	ExportPackage []JavaPackage
	Package       []JavaPackage
	KeyValues     []KeyValue
}

type MfIdentListener struct {
	currentKey string
	manifest   IgsoManifest

	parser.BaseManifestListener
}

func processStream(is antlr.CharStream) *parser.ManifestParser {
	lexer := parser.NewManifestLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return parser.NewManifestParser(stream)
}

func ProcessTsString(code string) *parser.ManifestParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

func Analysis(code string, fileName string) IgsoManifest {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Mf()

	listener := NewMfIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetResult()
}

func NewMfIdentListener(fileName string) *MfIdentListener {
	listener := &MfIdentListener{}
	listener.manifest = IgsoManifest{}
	return listener
}

func (s *MfIdentListener) EnterSection(ctx *parser.SectionContext) {
	if ctx.Key() != nil {
		s.currentKey = ctx.Key().GetText()
	}
}

func (s *MfIdentListener) EnterValue(ctx *parser.ValueContext) {
	s.manifest.KeyValues = append(s.manifest.KeyValues, KeyValue{
		Key:   s.currentKey,
		Value: ctx.GetText(),
	})

	if ctx.VERSION() != nil {
		versionText := ctx.STRING_LITERAL().GetText()
		versionInfo := versionText[2 : len(versionText)-2]
		split := strings.Split(versionInfo, ",")
		javaPackage := JavaPackage{
			Name:         ctx.OTHERS().GetText(),
			VersionInfo:  versionInfo,
			StartVersion: split[0],
			EndVersion:   split[1],
		}
		if len(ctx.AllConfigAssign()) > 0 {
			for _, config := range ctx.AllConfigAssign() {
				configAssign := (config).(*parser.ConfigAssignContext)
				javaPackage.Config = append(javaPackage.Config, KeyValue{
					Key: configAssign.AssignKey().GetText(),
					Value: configAssign.AssignValue().GetText(),
				})
			}
		}

		s.manifest.Package = append(s.manifest.Package, javaPackage)
	}
}

func (s *MfIdentListener) GetResult() IgsoManifest {
	return s.manifest
}
