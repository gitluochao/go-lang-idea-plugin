package main
var e = map[string]int{"key1":1, "key2":2}
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  VarDeclarationsImpl
    PsiElement(KEYWORD_VAR)('var')
    PsiWhiteSpace(' ')
    VarDeclarationImpl
      LiteralIdentifierImpl
        PsiElement(IDENTIFIER)('e')
      PsiWhiteSpace(' ')
      PsiElement(=)('=')
      PsiWhiteSpace(' ')
      LiteralExpressionImpl
        LiteralCompositeImpl
          TypeMapImpl
            PsiElement(KEYWORD_MAP)('map')
            PsiElement([)('[')
            TypeNameImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('string')
            PsiElement(])(']')
            TypeNameImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('int')
          LiteralCompositeValueImpl
            PsiElement({)('{')
            LiteralCompositeElementImpl
              CompositeLiteralElementKey
                LiteralExpressionImpl
                  LiteralStringImpl
                    PsiElement(LITERAL_STRING)('"key1"')
              PsiElement(:)(':')
              LiteralExpressionImpl
                LiteralIntegerImpl
                  PsiElement(LITERAL_INT)('1')
            PsiElement(,)(',')
            PsiWhiteSpace(' ')
            LiteralCompositeElementImpl
              CompositeLiteralElementKey
                LiteralExpressionImpl
                  LiteralStringImpl
                    PsiElement(LITERAL_STRING)('"key2"')
              PsiElement(:)(':')
              LiteralExpressionImpl
                LiteralIntegerImpl
                  PsiElement(LITERAL_INT)('2')
            PsiElement(})('}')
