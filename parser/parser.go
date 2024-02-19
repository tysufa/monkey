package parser

import (
  "qfa/ast"
  "qfa/lexer"
  "qfa/token"
)

type Parser struct{
  l *lexer.Lexer

  curToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser{
  p := &Parser{l: l}

  // read 2 tokens so curToken and peekToken are both set
  p.nextToken()
  p.nextToken()
}

func (p *Parser) nextToken(){
  p.curToken = p.peekToken
  p.peekToken = p.l.nextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
  return nil
}
