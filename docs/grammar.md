# Grammar

This is an attempt to document the grammar by James Mills and may fall behind as we update the syntax.
Please refer to the source code examples if you need the latest information.

```
(* Lexical basics *)
Identifier   = letter , { letter | digit | "_" } ;
IntegerLit   = DecInt | "0x" HexDigits | "0o" OctDigits | "0" | digit , {digit} ;
RuneLit      = "'" , ( escape | UnicodeScalar ) , "'" ;          (* see notes *)
StringLit    = '"' , { string_char | escape } , '"' ;
BoolLit      = "true" | "false" ;

letter       = "A"…"Z" | "a"…"z" ;
digit        = "0"…"9" ;
HexDigits    = hex , { hex } ;  hex = digit | "A"…"F" | "a"…"f" ;
OctDigits    = "0"…"7" ;

escape       = "\" , ( "0" | "t" | "n" | "r" | '"' | "'" | "\" ) ;
UnicodeScalar= any valid Unicode scalar value (examples include '世','界');

(* Program structure *)
CompilationUnit = { ImportDecl } , { TopLevelDecl } ;

ImportDecl   = "import" , Identifier ;                           (* e.g., import io, import os *)

TopLevelDecl = FuncDecl | ExternBlock ;

FuncDecl     = Identifier , "(" , [ ParamList ] , ")" ,
               [ "->" , ReturnTypes ] , Block ;

ParamList    = Param , { "," , Param } ;
Param        = Identifier , Type ;

ReturnTypes  = Type | "(" , NamedReturnList , ")" ;
NamedReturnList = NamedReturn , { "," , NamedReturn } ;
NamedReturn  = Identifier , Type ;

ExternBlock  = "extern" , "{" , { ExternNamespace } , "}" ;
ExternNamespace = Identifier , "{" , { ExternFuncSig } , "}" ;
ExternFuncSig = Identifier , "(" , [ ParamListExtern ] , ")" ,
                "->" , "(" , NamedReturnList , ")" ;
ParamListExtern= ParamExtern , { "," , ParamExtern } ;
ParamExtern  = Identifier , Type ;

(* Types *)
Type         = SimpleType | PtrType ;
SimpleType   = "int" | "uint" | "bool" | "byte" | "any" ;
PtrType      = "*" , Type ;                                      (* e.g., *byte, *any *)

(* Statements *)
Block        = "{" , { Statement } , "}" ;

Statement    = SimpleStmt
             | DeclStmt
             | IfStmt
             | LoopStmt
             | SwitchStmt
             | ReturnStmt ;

DeclStmt     = Identifier , ":=" , Expression ;                  (* infer type *)

SimpleStmt   = Assignment | ExprStmt | EmptyStmt ;

Assignment   = LValue , AssignOp , Expression ;
LValue       = Identifier | Selector ;                           (* basic lvalues *)
AssignOp     = "=" | "+=" | "-=" | "*=" | "/="
             | "%=" | "<<=" | ">>=" | "|=" | "&=" | "^=" ;

ExprStmt     = Expression ;
EmptyStmt    = ;

ReturnStmt   = "return" , [ Expression ] ;

IfStmt       = "if" , Expression , Block , [ "else" , Block ] ;

(* Looping: either infinite loop {…} or ranged iteration loop i := a..b {…} *)
LoopStmt     = "loop" , ( Block
                        | Identifier , ":=" , Expression , ".." , Expression , Block ) ;

(* Switch without a tag; a sequence of boolean guards with an optional default *)
SwitchStmt   = "switch" , "{" , { CaseClause } , [ DefaultClause ] , "}" ;
CaseClause   = Expression , "{" , { Statement } , "}" ;
DefaultClause= "_" , "{" , { Statement } , "}" ;

(* Expressions – precedence encoded by levels *)
Expression   = OrExpr ;

OrExpr       = AndExpr , { "||" , AndExpr } ;
AndExpr      = BitOrExpr , { "&&" , BitOrExpr } ;

BitOrExpr    = BitXorExpr , { "|" , BitXorExpr } ;
BitXorExpr   = BitAndExpr , { "^" , BitAndExpr } ;
BitAndExpr   = EqualityExpr , { "&" , EqualityExpr } ;

EqualityExpr = RelExpr , { ( "==" | "!=" ) , RelExpr } ;
RelExpr      = ShiftExpr , { ( "<" | "<=" | ">" | ">=" ) , ShiftExpr } ;

ShiftExpr    = AddExpr , { ( "<<" | ">>" ) , AddExpr } ;
AddExpr      = MulExpr , { ( "+" | "-" ) , MulExpr } ;
MulExpr      = UnaryExpr , { ( "*" | "/" | "%" ) , UnaryExpr } ;

UnaryExpr    = Primary | "-" , UnaryExpr | "+" , UnaryExpr | "!" , UnaryExpr ;

Primary      = Literal
             | Identifier
             | "(" , Expression , ")"
             | Call
             | Selector ;

Call         = Primary , "(" , [ ArgList ] , ")" ;
ArgList      = Expression , { "," , Expression } ;

Selector     = Primary , "." , Identifier ;                      (* e.g., message.ptr *)

Literal      = IntegerLit | StringLit | RuneLit | BoolLit ;

(* Notes:
   - Strings support property access “.ptr” in examples to pass pointers to externs.
   - Range syntax is “a..b” in loop headers and is inclusive in examples that enumerate 2..100.
   - Default case in switch is “_”.
   - There is no ‘break’ shown; loops exit via ‘return’ in examples/tests. *)
```