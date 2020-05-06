// Code generated from SysVerilogHDL.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SysVerilogHDL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSysVerilogHDLListener is a complete listener for a parse tree produced by SysVerilogHDLParser.
type BaseSysVerilogHDLListener struct{}

var _ SysVerilogHDLListener = &BaseSysVerilogHDLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSysVerilogHDLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSysVerilogHDLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSysVerilogHDLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSysVerilogHDLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModule_keyword is called when production module_keyword is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_keyword(ctx *Module_keywordContext) {}

// ExitModule_keyword is called when production module_keyword is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_keyword(ctx *Module_keywordContext) {}

// EnterStruct_keyword is called when production struct_keyword is entered.
func (s *BaseSysVerilogHDLListener) EnterStruct_keyword(ctx *Struct_keywordContext) {}

// ExitStruct_keyword is called when production struct_keyword is exited.
func (s *BaseSysVerilogHDLListener) ExitStruct_keyword(ctx *Struct_keywordContext) {}

// EnterAny_case_keyword is called when production any_case_keyword is entered.
func (s *BaseSysVerilogHDLListener) EnterAny_case_keyword(ctx *Any_case_keywordContext) {}

// ExitAny_case_keyword is called when production any_case_keyword is exited.
func (s *BaseSysVerilogHDLListener) ExitAny_case_keyword(ctx *Any_case_keywordContext) {}

// EnterSemicolon is called when production semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterSemicolon(ctx *SemicolonContext) {}

// ExitSemicolon is called when production semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitSemicolon(ctx *SemicolonContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseSysVerilogHDLListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseSysVerilogHDLListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterBinary_operator is called when production binary_operator is entered.
func (s *BaseSysVerilogHDLListener) EnterBinary_operator(ctx *Binary_operatorContext) {}

// ExitBinary_operator is called when production binary_operator is exited.
func (s *BaseSysVerilogHDLListener) ExitBinary_operator(ctx *Binary_operatorContext) {}

// EnterUnary_assign_operator is called when production unary_assign_operator is entered.
func (s *BaseSysVerilogHDLListener) EnterUnary_assign_operator(ctx *Unary_assign_operatorContext) {}

// ExitUnary_assign_operator is called when production unary_assign_operator is exited.
func (s *BaseSysVerilogHDLListener) ExitUnary_assign_operator(ctx *Unary_assign_operatorContext) {}

// EnterBinary_assign_operator is called when production binary_assign_operator is entered.
func (s *BaseSysVerilogHDLListener) EnterBinary_assign_operator(ctx *Binary_assign_operatorContext) {}

// ExitBinary_assign_operator is called when production binary_assign_operator is exited.
func (s *BaseSysVerilogHDLListener) ExitBinary_assign_operator(ctx *Binary_assign_operatorContext) {}

// EnterSource_text is called when production source_text is entered.
func (s *BaseSysVerilogHDLListener) EnterSource_text(ctx *Source_textContext) {}

// ExitSource_text is called when production source_text is exited.
func (s *BaseSysVerilogHDLListener) ExitSource_text(ctx *Source_textContext) {}

// EnterDescription_star is called when production description_star is entered.
func (s *BaseSysVerilogHDLListener) EnterDescription_star(ctx *Description_starContext) {}

// ExitDescription_star is called when production description_star is exited.
func (s *BaseSysVerilogHDLListener) ExitDescription_star(ctx *Description_starContext) {}

// EnterHeader_text is called when production header_text is entered.
func (s *BaseSysVerilogHDLListener) EnterHeader_text(ctx *Header_textContext) {}

// ExitHeader_text is called when production header_text is exited.
func (s *BaseSysVerilogHDLListener) ExitHeader_text(ctx *Header_textContext) {}

// EnterDesign_attribute is called when production design_attribute is entered.
func (s *BaseSysVerilogHDLListener) EnterDesign_attribute(ctx *Design_attributeContext) {}

// ExitDesign_attribute is called when production design_attribute is exited.
func (s *BaseSysVerilogHDLListener) ExitDesign_attribute(ctx *Design_attributeContext) {}

// EnterCompiler_directive is called when production compiler_directive is entered.
func (s *BaseSysVerilogHDLListener) EnterCompiler_directive(ctx *Compiler_directiveContext) {}

// ExitCompiler_directive is called when production compiler_directive is exited.
func (s *BaseSysVerilogHDLListener) ExitCompiler_directive(ctx *Compiler_directiveContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseSysVerilogHDLListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseSysVerilogHDLListener) ExitDescription(ctx *DescriptionContext) {}

// EnterModule_declaration is called when production module_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_declaration(ctx *Module_declarationContext) {}

// ExitModule_declaration is called when production module_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_declaration(ctx *Module_declarationContext) {}

// EnterModule_identifier is called when production module_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_identifier(ctx *Module_identifierContext) {}

// ExitModule_identifier is called when production module_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_identifier(ctx *Module_identifierContext) {}

// EnterModule_interface is called when production module_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_interface(ctx *Module_interfaceContext) {}

// ExitModule_interface is called when production module_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_interface(ctx *Module_interfaceContext) {}

// EnterModule_parameter_interface is called when production module_parameter_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_parameter_interface(ctx *Module_parameter_interfaceContext) {
}

// ExitModule_parameter_interface is called when production module_parameter_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_parameter_interface(ctx *Module_parameter_interfaceContext) {
}

// EnterModule_port_interface is called when production module_port_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_port_interface(ctx *Module_port_interfaceContext) {}

// ExitModule_port_interface is called when production module_port_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_port_interface(ctx *Module_port_interfaceContext) {}

// EnterModule_item_star is called when production module_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_item_star(ctx *Module_item_starContext) {}

// ExitModule_item_star is called when production module_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_item_star(ctx *Module_item_starContext) {}

// EnterModule_item is called when production module_item is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_item(ctx *Module_itemContext) {}

// ExitModule_item is called when production module_item is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_item(ctx *Module_itemContext) {}

// EnterColon_module_identifier is called when production colon_module_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterColon_module_identifier(ctx *Colon_module_identifierContext) {
}

// ExitColon_module_identifier is called when production colon_module_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitColon_module_identifier(ctx *Colon_module_identifierContext) {}

// EnterPackage_declaration is called when production package_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterPackage_declaration(ctx *Package_declarationContext) {}

// ExitPackage_declaration is called when production package_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitPackage_declaration(ctx *Package_declarationContext) {}

// EnterPackage_identifier is called when production package_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterPackage_identifier(ctx *Package_identifierContext) {}

// ExitPackage_identifier is called when production package_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitPackage_identifier(ctx *Package_identifierContext) {}

// EnterColon_package_identifier is called when production colon_package_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterColon_package_identifier(ctx *Colon_package_identifierContext) {
}

// ExitColon_package_identifier is called when production colon_package_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitColon_package_identifier(ctx *Colon_package_identifierContext) {
}

// EnterPackage_item_star is called when production package_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterPackage_item_star(ctx *Package_item_starContext) {}

// ExitPackage_item_star is called when production package_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitPackage_item_star(ctx *Package_item_starContext) {}

// EnterPackage_item is called when production package_item is entered.
func (s *BaseSysVerilogHDLListener) EnterPackage_item(ctx *Package_itemContext) {}

// ExitPackage_item is called when production package_item is exited.
func (s *BaseSysVerilogHDLListener) ExitPackage_item(ctx *Package_itemContext) {}

// EnterImport_package is called when production import_package is entered.
func (s *BaseSysVerilogHDLListener) EnterImport_package(ctx *Import_packageContext) {}

// ExitImport_package is called when production import_package is exited.
func (s *BaseSysVerilogHDLListener) ExitImport_package(ctx *Import_packageContext) {}

// EnterPackage_item_identifier is called when production package_item_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterPackage_item_identifier(ctx *Package_item_identifierContext) {
}

// ExitPackage_item_identifier is called when production package_item_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitPackage_item_identifier(ctx *Package_item_identifierContext) {}

// EnterParameter_item_semicolon is called when production parameter_item_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterParameter_item_semicolon(ctx *Parameter_item_semicolonContext) {
}

// ExitParameter_item_semicolon is called when production parameter_item_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitParameter_item_semicolon(ctx *Parameter_item_semicolonContext) {
}

// EnterParameter_item is called when production parameter_item is entered.
func (s *BaseSysVerilogHDLListener) EnterParameter_item(ctx *Parameter_itemContext) {}

// ExitParameter_item is called when production parameter_item is exited.
func (s *BaseSysVerilogHDLListener) ExitParameter_item(ctx *Parameter_itemContext) {}

// EnterAttr_port_item_semicolon is called when production attr_port_item_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_port_item_semicolon(ctx *Attr_port_item_semicolonContext) {
}

// ExitAttr_port_item_semicolon is called when production attr_port_item_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_port_item_semicolon(ctx *Attr_port_item_semicolonContext) {
}

// EnterAttr_variable_item_semicolon is called when production attr_variable_item_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_variable_item_semicolon(ctx *Attr_variable_item_semicolonContext) {
}

// ExitAttr_variable_item_semicolon is called when production attr_variable_item_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_variable_item_semicolon(ctx *Attr_variable_item_semicolonContext) {
}

// EnterVariable_item is called when production variable_item is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_item(ctx *Variable_itemContext) {}

// ExitVariable_item is called when production variable_item is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_item(ctx *Variable_itemContext) {}

// EnterSubroutine_item_semicolon is called when production subroutine_item_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterSubroutine_item_semicolon(ctx *Subroutine_item_semicolonContext) {
}

// ExitSubroutine_item_semicolon is called when production subroutine_item_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitSubroutine_item_semicolon(ctx *Subroutine_item_semicolonContext) {
}

// EnterSubroutine_item is called when production subroutine_item is entered.
func (s *BaseSysVerilogHDLListener) EnterSubroutine_item(ctx *Subroutine_itemContext) {}

// ExitSubroutine_item is called when production subroutine_item is exited.
func (s *BaseSysVerilogHDLListener) ExitSubroutine_item(ctx *Subroutine_itemContext) {}

// EnterAttr_construct_item is called when production attr_construct_item is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_construct_item(ctx *Attr_construct_itemContext) {}

// ExitAttr_construct_item is called when production attr_construct_item is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_construct_item(ctx *Attr_construct_itemContext) {}

// EnterConstruct_item is called when production construct_item is entered.
func (s *BaseSysVerilogHDLListener) EnterConstruct_item(ctx *Construct_itemContext) {}

// ExitConstruct_item is called when production construct_item is exited.
func (s *BaseSysVerilogHDLListener) ExitConstruct_item(ctx *Construct_itemContext) {}

// EnterAttr_component_item is called when production attr_component_item is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_component_item(ctx *Attr_component_itemContext) {}

// ExitAttr_component_item is called when production attr_component_item is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_component_item(ctx *Attr_component_itemContext) {}

// EnterComponent_item is called when production component_item is entered.
func (s *BaseSysVerilogHDLListener) EnterComponent_item(ctx *Component_itemContext) {}

// ExitComponent_item is called when production component_item is exited.
func (s *BaseSysVerilogHDLListener) ExitComponent_item(ctx *Component_itemContext) {}

// EnterCompiler_item is called when production compiler_item is entered.
func (s *BaseSysVerilogHDLListener) EnterCompiler_item(ctx *Compiler_itemContext) {}

// ExitCompiler_item is called when production compiler_item is exited.
func (s *BaseSysVerilogHDLListener) ExitCompiler_item(ctx *Compiler_itemContext) {}

// EnterType_item is called when production type_item is entered.
func (s *BaseSysVerilogHDLListener) EnterType_item(ctx *Type_itemContext) {}

// ExitType_item is called when production type_item is exited.
func (s *BaseSysVerilogHDLListener) ExitType_item(ctx *Type_itemContext) {}

// EnterNull_item is called when production null_item is entered.
func (s *BaseSysVerilogHDLListener) EnterNull_item(ctx *Null_itemContext) {}

// ExitNull_item is called when production null_item is exited.
func (s *BaseSysVerilogHDLListener) ExitNull_item(ctx *Null_itemContext) {}

// EnterList_of_interface_parameters is called when production list_of_interface_parameters is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_interface_parameters(ctx *List_of_interface_parametersContext) {
}

// ExitList_of_interface_parameters is called when production list_of_interface_parameters is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_interface_parameters(ctx *List_of_interface_parametersContext) {
}

// EnterList_of_parameter_declarations is called when production list_of_parameter_declarations is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_parameter_declarations(ctx *List_of_parameter_declarationsContext) {
}

// ExitList_of_parameter_declarations is called when production list_of_parameter_declarations is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_parameter_declarations(ctx *List_of_parameter_declarationsContext) {
}

// EnterComma_parameter_declaration_star is called when production comma_parameter_declaration_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_parameter_declaration_star(ctx *Comma_parameter_declaration_starContext) {
}

// ExitComma_parameter_declaration_star is called when production comma_parameter_declaration_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_parameter_declaration_star(ctx *Comma_parameter_declaration_starContext) {
}

// EnterComma_parameter_declaration is called when production comma_parameter_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_parameter_declaration(ctx *Comma_parameter_declarationContext) {
}

// ExitComma_parameter_declaration is called when production comma_parameter_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_parameter_declaration(ctx *Comma_parameter_declarationContext) {
}

// EnterList_of_parameter_descriptions is called when production list_of_parameter_descriptions is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_parameter_descriptions(ctx *List_of_parameter_descriptionsContext) {
}

// ExitList_of_parameter_descriptions is called when production list_of_parameter_descriptions is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_parameter_descriptions(ctx *List_of_parameter_descriptionsContext) {
}

// EnterParam_declaration is called when production param_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterParam_declaration(ctx *Param_declarationContext) {}

// ExitParam_declaration is called when production param_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitParam_declaration(ctx *Param_declarationContext) {}

// EnterParam_description is called when production param_description is entered.
func (s *BaseSysVerilogHDLListener) EnterParam_description(ctx *Param_descriptionContext) {}

// ExitParam_description is called when production param_description is exited.
func (s *BaseSysVerilogHDLListener) ExitParam_description(ctx *Param_descriptionContext) {}

// EnterParameter_declaration is called when production parameter_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterParameter_declaration(ctx *Parameter_declarationContext) {}

// ExitParameter_declaration is called when production parameter_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitParameter_declaration(ctx *Parameter_declarationContext) {}

// EnterLocal_parameter_declaration is called when production local_parameter_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterLocal_parameter_declaration(ctx *Local_parameter_declarationContext) {
}

// ExitLocal_parameter_declaration is called when production local_parameter_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitLocal_parameter_declaration(ctx *Local_parameter_declarationContext) {
}

// EnterParameter_override is called when production parameter_override is entered.
func (s *BaseSysVerilogHDLListener) EnterParameter_override(ctx *Parameter_overrideContext) {}

// ExitParameter_override is called when production parameter_override is exited.
func (s *BaseSysVerilogHDLListener) ExitParameter_override(ctx *Parameter_overrideContext) {}

// EnterList_of_tf_interface_ports is called when production list_of_tf_interface_ports is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_tf_interface_ports(ctx *List_of_tf_interface_portsContext) {
}

// ExitList_of_tf_interface_ports is called when production list_of_tf_interface_ports is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_tf_interface_ports(ctx *List_of_tf_interface_portsContext) {
}

// EnterList_of_tf_port_declarations is called when production list_of_tf_port_declarations is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_tf_port_declarations(ctx *List_of_tf_port_declarationsContext) {
}

// ExitList_of_tf_port_declarations is called when production list_of_tf_port_declarations is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_tf_port_declarations(ctx *List_of_tf_port_declarationsContext) {
}

// EnterList_of_tf_port_declarations_comma is called when production list_of_tf_port_declarations_comma is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_tf_port_declarations_comma(ctx *List_of_tf_port_declarations_commaContext) {
}

// ExitList_of_tf_port_declarations_comma is called when production list_of_tf_port_declarations_comma is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_tf_port_declarations_comma(ctx *List_of_tf_port_declarations_commaContext) {
}

// EnterComma_attr_tf_port_declaration_star is called when production comma_attr_tf_port_declaration_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_attr_tf_port_declaration_star(ctx *Comma_attr_tf_port_declaration_starContext) {
}

// ExitComma_attr_tf_port_declaration_star is called when production comma_attr_tf_port_declaration_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_attr_tf_port_declaration_star(ctx *Comma_attr_tf_port_declaration_starContext) {
}

// EnterComma_attr_tf_port_declaration is called when production comma_attr_tf_port_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_attr_tf_port_declaration(ctx *Comma_attr_tf_port_declarationContext) {
}

// ExitComma_attr_tf_port_declaration is called when production comma_attr_tf_port_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_attr_tf_port_declaration(ctx *Comma_attr_tf_port_declarationContext) {
}

// EnterList_of_tf_port_declarations_semicolon is called when production list_of_tf_port_declarations_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_tf_port_declarations_semicolon(ctx *List_of_tf_port_declarations_semicolonContext) {
}

// ExitList_of_tf_port_declarations_semicolon is called when production list_of_tf_port_declarations_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_tf_port_declarations_semicolon(ctx *List_of_tf_port_declarations_semicolonContext) {
}

// EnterAttr_tf_port_declaration_semicolon_plus is called when production attr_tf_port_declaration_semicolon_plus is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_tf_port_declaration_semicolon_plus(ctx *Attr_tf_port_declaration_semicolon_plusContext) {
}

// ExitAttr_tf_port_declaration_semicolon_plus is called when production attr_tf_port_declaration_semicolon_plus is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_tf_port_declaration_semicolon_plus(ctx *Attr_tf_port_declaration_semicolon_plusContext) {
}

// EnterAttr_tf_port_declaration_semicolon_star is called when production attr_tf_port_declaration_semicolon_star is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_tf_port_declaration_semicolon_star(ctx *Attr_tf_port_declaration_semicolon_starContext) {
}

// ExitAttr_tf_port_declaration_semicolon_star is called when production attr_tf_port_declaration_semicolon_star is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_tf_port_declaration_semicolon_star(ctx *Attr_tf_port_declaration_semicolon_starContext) {
}

// EnterAttr_tf_port_declaration_semicolon is called when production attr_tf_port_declaration_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_tf_port_declaration_semicolon(ctx *Attr_tf_port_declaration_semicolonContext) {
}

// ExitAttr_tf_port_declaration_semicolon is called when production attr_tf_port_declaration_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_tf_port_declaration_semicolon(ctx *Attr_tf_port_declaration_semicolonContext) {
}

// EnterAttr_tf_port_declaration is called when production attr_tf_port_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_tf_port_declaration(ctx *Attr_tf_port_declarationContext) {
}

// ExitAttr_tf_port_declaration is called when production attr_tf_port_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_tf_port_declaration(ctx *Attr_tf_port_declarationContext) {
}

// EnterTf_port_declaration is called when production tf_port_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterTf_port_declaration(ctx *Tf_port_declarationContext) {}

// ExitTf_port_declaration is called when production tf_port_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitTf_port_declaration(ctx *Tf_port_declarationContext) {}

// EnterList_of_interface_ports is called when production list_of_interface_ports is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_interface_ports(ctx *List_of_interface_portsContext) {
}

// ExitList_of_interface_ports is called when production list_of_interface_ports is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_interface_ports(ctx *List_of_interface_portsContext) {}

// EnterList_of_port_identifiers is called when production list_of_port_identifiers is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_port_identifiers(ctx *List_of_port_identifiersContext) {
}

// ExitList_of_port_identifiers is called when production list_of_port_identifiers is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_port_identifiers(ctx *List_of_port_identifiersContext) {
}

// EnterComma_port_identifier_star is called when production comma_port_identifier_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_port_identifier_star(ctx *Comma_port_identifier_starContext) {
}

// ExitComma_port_identifier_star is called when production comma_port_identifier_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_port_identifier_star(ctx *Comma_port_identifier_starContext) {
}

// EnterComma_port_identifier is called when production comma_port_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_port_identifier(ctx *Comma_port_identifierContext) {}

// ExitComma_port_identifier is called when production comma_port_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_port_identifier(ctx *Comma_port_identifierContext) {}

// EnterPort_identifier is called when production port_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterPort_identifier(ctx *Port_identifierContext) {}

// ExitPort_identifier is called when production port_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitPort_identifier(ctx *Port_identifierContext) {}

// EnterList_of_port_declarations is called when production list_of_port_declarations is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_port_declarations(ctx *List_of_port_declarationsContext) {
}

// ExitList_of_port_declarations is called when production list_of_port_declarations is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_port_declarations(ctx *List_of_port_declarationsContext) {
}

// EnterList_of_port_declarations_comma is called when production list_of_port_declarations_comma is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_port_declarations_comma(ctx *List_of_port_declarations_commaContext) {
}

// ExitList_of_port_declarations_comma is called when production list_of_port_declarations_comma is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_port_declarations_comma(ctx *List_of_port_declarations_commaContext) {
}

// EnterComma_attr_port_declaration_star is called when production comma_attr_port_declaration_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_attr_port_declaration_star(ctx *Comma_attr_port_declaration_starContext) {
}

// ExitComma_attr_port_declaration_star is called when production comma_attr_port_declaration_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_attr_port_declaration_star(ctx *Comma_attr_port_declaration_starContext) {
}

// EnterComma_attr_port_declaration is called when production comma_attr_port_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_attr_port_declaration(ctx *Comma_attr_port_declarationContext) {
}

// ExitComma_attr_port_declaration is called when production comma_attr_port_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_attr_port_declaration(ctx *Comma_attr_port_declarationContext) {
}

// EnterList_of_port_declarations_semicolon is called when production list_of_port_declarations_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_port_declarations_semicolon(ctx *List_of_port_declarations_semicolonContext) {
}

// ExitList_of_port_declarations_semicolon is called when production list_of_port_declarations_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_port_declarations_semicolon(ctx *List_of_port_declarations_semicolonContext) {
}

// EnterAttr_port_declaration_semicolon_plus is called when production attr_port_declaration_semicolon_plus is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_port_declaration_semicolon_plus(ctx *Attr_port_declaration_semicolon_plusContext) {
}

// ExitAttr_port_declaration_semicolon_plus is called when production attr_port_declaration_semicolon_plus is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_port_declaration_semicolon_plus(ctx *Attr_port_declaration_semicolon_plusContext) {
}

// EnterAttr_port_declaration_semicolon_star is called when production attr_port_declaration_semicolon_star is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_port_declaration_semicolon_star(ctx *Attr_port_declaration_semicolon_starContext) {
}

// ExitAttr_port_declaration_semicolon_star is called when production attr_port_declaration_semicolon_star is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_port_declaration_semicolon_star(ctx *Attr_port_declaration_semicolon_starContext) {
}

// EnterAttr_port_declaration_semicolon is called when production attr_port_declaration_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_port_declaration_semicolon(ctx *Attr_port_declaration_semicolonContext) {
}

// ExitAttr_port_declaration_semicolon is called when production attr_port_declaration_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_port_declaration_semicolon(ctx *Attr_port_declaration_semicolonContext) {
}

// EnterAttr_port_declaration is called when production attr_port_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_port_declaration(ctx *Attr_port_declarationContext) {}

// ExitAttr_port_declaration is called when production attr_port_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_port_declaration(ctx *Attr_port_declarationContext) {}

// EnterPort_declaration is called when production port_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterPort_declaration(ctx *Port_declarationContext) {}

// ExitPort_declaration is called when production port_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitPort_declaration(ctx *Port_declarationContext) {}

// EnterPort_description is called when production port_description is entered.
func (s *BaseSysVerilogHDLListener) EnterPort_description(ctx *Port_descriptionContext) {}

// ExitPort_description is called when production port_description is exited.
func (s *BaseSysVerilogHDLListener) ExitPort_description(ctx *Port_descriptionContext) {}

// EnterInout_description is called when production inout_description is entered.
func (s *BaseSysVerilogHDLListener) EnterInout_description(ctx *Inout_descriptionContext) {}

// ExitInout_description is called when production inout_description is exited.
func (s *BaseSysVerilogHDLListener) ExitInout_description(ctx *Inout_descriptionContext) {}

// EnterInput_description is called when production input_description is entered.
func (s *BaseSysVerilogHDLListener) EnterInput_description(ctx *Input_descriptionContext) {}

// ExitInput_description is called when production input_description is exited.
func (s *BaseSysVerilogHDLListener) ExitInput_description(ctx *Input_descriptionContext) {}

// EnterOutput_description is called when production output_description is entered.
func (s *BaseSysVerilogHDLListener) EnterOutput_description(ctx *Output_descriptionContext) {}

// ExitOutput_description is called when production output_description is exited.
func (s *BaseSysVerilogHDLListener) ExitOutput_description(ctx *Output_descriptionContext) {}

// EnterRef_description is called when production ref_description is entered.
func (s *BaseSysVerilogHDLListener) EnterRef_description(ctx *Ref_descriptionContext) {}

// ExitRef_description is called when production ref_description is exited.
func (s *BaseSysVerilogHDLListener) ExitRef_description(ctx *Ref_descriptionContext) {}

// EnterTf_declaration is called when production tf_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterTf_declaration(ctx *Tf_declarationContext) {}

// ExitTf_declaration is called when production tf_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitTf_declaration(ctx *Tf_declarationContext) {}

// EnterInout_declaration is called when production inout_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterInout_declaration(ctx *Inout_declarationContext) {}

// ExitInout_declaration is called when production inout_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitInout_declaration(ctx *Inout_declarationContext) {}

// EnterInput_declaration is called when production input_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterInput_declaration(ctx *Input_declarationContext) {}

// ExitInput_declaration is called when production input_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitInput_declaration(ctx *Input_declarationContext) {}

// EnterOutput_declaration is called when production output_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterOutput_declaration(ctx *Output_declarationContext) {}

// ExitOutput_declaration is called when production output_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitOutput_declaration(ctx *Output_declarationContext) {}

// EnterRef_declaration is called when production ref_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterRef_declaration(ctx *Ref_declarationContext) {}

// ExitRef_declaration is called when production ref_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitRef_declaration(ctx *Ref_declarationContext) {}

// EnterUser_type is called when production user_type is entered.
func (s *BaseSysVerilogHDLListener) EnterUser_type(ctx *User_typeContext) {}

// ExitUser_type is called when production user_type is exited.
func (s *BaseSysVerilogHDLListener) ExitUser_type(ctx *User_typeContext) {}

// EnterUser_type_identifer is called when production user_type_identifer is entered.
func (s *BaseSysVerilogHDLListener) EnterUser_type_identifer(ctx *User_type_identiferContext) {}

// ExitUser_type_identifer is called when production user_type_identifer is exited.
func (s *BaseSysVerilogHDLListener) ExitUser_type_identifer(ctx *User_type_identiferContext) {}

// EnterDimension_plus is called when production dimension_plus is entered.
func (s *BaseSysVerilogHDLListener) EnterDimension_plus(ctx *Dimension_plusContext) {}

// ExitDimension_plus is called when production dimension_plus is exited.
func (s *BaseSysVerilogHDLListener) ExitDimension_plus(ctx *Dimension_plusContext) {}

// EnterDimension_star is called when production dimension_star is entered.
func (s *BaseSysVerilogHDLListener) EnterDimension_star(ctx *Dimension_starContext) {}

// ExitDimension_star is called when production dimension_star is exited.
func (s *BaseSysVerilogHDLListener) ExitDimension_star(ctx *Dimension_starContext) {}

// EnterDimension is called when production dimension is entered.
func (s *BaseSysVerilogHDLListener) EnterDimension(ctx *DimensionContext) {}

// ExitDimension is called when production dimension is exited.
func (s *BaseSysVerilogHDLListener) ExitDimension(ctx *DimensionContext) {}

// EnterRange_expression is called when production range_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterRange_expression(ctx *Range_expressionContext) {}

// ExitRange_expression is called when production range_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitRange_expression(ctx *Range_expressionContext) {}

// EnterIndex_expression is called when production index_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterIndex_expression(ctx *Index_expressionContext) {}

// ExitIndex_expression is called when production index_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitIndex_expression(ctx *Index_expressionContext) {}

// EnterSb_range is called when production sb_range is entered.
func (s *BaseSysVerilogHDLListener) EnterSb_range(ctx *Sb_rangeContext) {}

// ExitSb_range is called when production sb_range is exited.
func (s *BaseSysVerilogHDLListener) ExitSb_range(ctx *Sb_rangeContext) {}

// EnterBase_increment_range is called when production base_increment_range is entered.
func (s *BaseSysVerilogHDLListener) EnterBase_increment_range(ctx *Base_increment_rangeContext) {}

// ExitBase_increment_range is called when production base_increment_range is exited.
func (s *BaseSysVerilogHDLListener) ExitBase_increment_range(ctx *Base_increment_rangeContext) {}

// EnterBase_decrement_range is called when production base_decrement_range is entered.
func (s *BaseSysVerilogHDLListener) EnterBase_decrement_range(ctx *Base_decrement_rangeContext) {}

// ExitBase_decrement_range is called when production base_decrement_range is exited.
func (s *BaseSysVerilogHDLListener) ExitBase_decrement_range(ctx *Base_decrement_rangeContext) {}

// EnterBase_expression is called when production base_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterBase_expression(ctx *Base_expressionContext) {}

// ExitBase_expression is called when production base_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitBase_expression(ctx *Base_expressionContext) {}

// EnterNet_type is called when production net_type is entered.
func (s *BaseSysVerilogHDLListener) EnterNet_type(ctx *Net_typeContext) {}

// ExitNet_type is called when production net_type is exited.
func (s *BaseSysVerilogHDLListener) ExitNet_type(ctx *Net_typeContext) {}

// EnterDrive_strength is called when production drive_strength is entered.
func (s *BaseSysVerilogHDLListener) EnterDrive_strength(ctx *Drive_strengthContext) {}

// ExitDrive_strength is called when production drive_strength is exited.
func (s *BaseSysVerilogHDLListener) ExitDrive_strength(ctx *Drive_strengthContext) {}

// EnterDrive_strength_value_0 is called when production drive_strength_value_0 is entered.
func (s *BaseSysVerilogHDLListener) EnterDrive_strength_value_0(ctx *Drive_strength_value_0Context) {}

// ExitDrive_strength_value_0 is called when production drive_strength_value_0 is exited.
func (s *BaseSysVerilogHDLListener) ExitDrive_strength_value_0(ctx *Drive_strength_value_0Context) {}

// EnterDrive_strength_value_1 is called when production drive_strength_value_1 is entered.
func (s *BaseSysVerilogHDLListener) EnterDrive_strength_value_1(ctx *Drive_strength_value_1Context) {}

// ExitDrive_strength_value_1 is called when production drive_strength_value_1 is exited.
func (s *BaseSysVerilogHDLListener) ExitDrive_strength_value_1(ctx *Drive_strength_value_1Context) {}

// EnterStrength0 is called when production strength0 is entered.
func (s *BaseSysVerilogHDLListener) EnterStrength0(ctx *Strength0Context) {}

// ExitStrength0 is called when production strength0 is exited.
func (s *BaseSysVerilogHDLListener) ExitStrength0(ctx *Strength0Context) {}

// EnterStrength1 is called when production strength1 is entered.
func (s *BaseSysVerilogHDLListener) EnterStrength1(ctx *Strength1Context) {}

// ExitStrength1 is called when production strength1 is exited.
func (s *BaseSysVerilogHDLListener) ExitStrength1(ctx *Strength1Context) {}

// EnterHighz0 is called when production highz0 is entered.
func (s *BaseSysVerilogHDLListener) EnterHighz0(ctx *Highz0Context) {}

// ExitHighz0 is called when production highz0 is exited.
func (s *BaseSysVerilogHDLListener) ExitHighz0(ctx *Highz0Context) {}

// EnterHighz1 is called when production highz1 is entered.
func (s *BaseSysVerilogHDLListener) EnterHighz1(ctx *Highz1Context) {}

// ExitHighz1 is called when production highz1 is exited.
func (s *BaseSysVerilogHDLListener) ExitHighz1(ctx *Highz1Context) {}

// EnterCharge_strength is called when production charge_strength is entered.
func (s *BaseSysVerilogHDLListener) EnterCharge_strength(ctx *Charge_strengthContext) {}

// ExitCharge_strength is called when production charge_strength is exited.
func (s *BaseSysVerilogHDLListener) ExitCharge_strength(ctx *Charge_strengthContext) {}

// EnterCharge_size is called when production charge_size is entered.
func (s *BaseSysVerilogHDLListener) EnterCharge_size(ctx *Charge_sizeContext) {}

// ExitCharge_size is called when production charge_size is exited.
func (s *BaseSysVerilogHDLListener) ExitCharge_size(ctx *Charge_sizeContext) {}

// EnterList_of_variable_descriptions is called when production list_of_variable_descriptions is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_variable_descriptions(ctx *List_of_variable_descriptionsContext) {
}

// ExitList_of_variable_descriptions is called when production list_of_variable_descriptions is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_variable_descriptions(ctx *List_of_variable_descriptionsContext) {
}

// EnterComma_variable_description_star is called when production comma_variable_description_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_variable_description_star(ctx *Comma_variable_description_starContext) {
}

// ExitComma_variable_description_star is called when production comma_variable_description_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_variable_description_star(ctx *Comma_variable_description_starContext) {
}

// EnterComma_variable_description is called when production comma_variable_description is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_variable_description(ctx *Comma_variable_descriptionContext) {
}

// ExitComma_variable_description is called when production comma_variable_description is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_variable_description(ctx *Comma_variable_descriptionContext) {
}

// EnterVariable_description is called when production variable_description is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_description(ctx *Variable_descriptionContext) {}

// ExitVariable_description is called when production variable_description is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_description(ctx *Variable_descriptionContext) {}

// EnterVariable_identifier is called when production variable_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_identifier(ctx *Variable_identifierContext) {}

// ExitVariable_identifier is called when production variable_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_identifier(ctx *Variable_identifierContext) {}

// EnterList_of_hierarchical_variable_descriptions is called when production list_of_hierarchical_variable_descriptions is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_hierarchical_variable_descriptions(ctx *List_of_hierarchical_variable_descriptionsContext) {
}

// ExitList_of_hierarchical_variable_descriptions is called when production list_of_hierarchical_variable_descriptions is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_hierarchical_variable_descriptions(ctx *List_of_hierarchical_variable_descriptionsContext) {
}

// EnterComma_hierarchical_variable_description_star is called when production comma_hierarchical_variable_description_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_hierarchical_variable_description_star(ctx *Comma_hierarchical_variable_description_starContext) {
}

// ExitComma_hierarchical_variable_description_star is called when production comma_hierarchical_variable_description_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_hierarchical_variable_description_star(ctx *Comma_hierarchical_variable_description_starContext) {
}

// EnterComma_hierarchical_variable_description is called when production comma_hierarchical_variable_description is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_hierarchical_variable_description(ctx *Comma_hierarchical_variable_descriptionContext) {
}

// ExitComma_hierarchical_variable_description is called when production comma_hierarchical_variable_description is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_hierarchical_variable_description(ctx *Comma_hierarchical_variable_descriptionContext) {
}

// EnterHierarchical_variable_description is called when production hierarchical_variable_description is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_variable_description(ctx *Hierarchical_variable_descriptionContext) {
}

// ExitHierarchical_variable_description is called when production hierarchical_variable_description is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_variable_description(ctx *Hierarchical_variable_descriptionContext) {
}

// EnterHierarchical_variable_identifier is called when production hierarchical_variable_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_variable_identifier(ctx *Hierarchical_variable_identifierContext) {
}

// ExitHierarchical_variable_identifier is called when production hierarchical_variable_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_variable_identifier(ctx *Hierarchical_variable_identifierContext) {
}

// EnterNet_declaration is called when production net_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterNet_declaration(ctx *Net_declarationContext) {}

// ExitNet_declaration is called when production net_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitNet_declaration(ctx *Net_declarationContext) {}

// EnterReg_declaration is called when production reg_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterReg_declaration(ctx *Reg_declarationContext) {}

// ExitReg_declaration is called when production reg_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitReg_declaration(ctx *Reg_declarationContext) {}

// EnterLogic_declaration is called when production logic_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterLogic_declaration(ctx *Logic_declarationContext) {}

// ExitLogic_declaration is called when production logic_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitLogic_declaration(ctx *Logic_declarationContext) {}

// EnterBits_type is called when production bits_type is entered.
func (s *BaseSysVerilogHDLListener) EnterBits_type(ctx *Bits_typeContext) {}

// ExitBits_type is called when production bits_type is exited.
func (s *BaseSysVerilogHDLListener) ExitBits_type(ctx *Bits_typeContext) {}

// EnterBits_declaration is called when production bits_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterBits_declaration(ctx *Bits_declarationContext) {}

// ExitBits_declaration is called when production bits_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitBits_declaration(ctx *Bits_declarationContext) {}

// EnterInteger_declaration is called when production integer_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterInteger_declaration(ctx *Integer_declarationContext) {}

// ExitInteger_declaration is called when production integer_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitInteger_declaration(ctx *Integer_declarationContext) {}

// EnterInt_declaration is called when production int_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterInt_declaration(ctx *Int_declarationContext) {}

// ExitInt_declaration is called when production int_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitInt_declaration(ctx *Int_declarationContext) {}

// EnterReal_declaration is called when production real_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterReal_declaration(ctx *Real_declarationContext) {}

// ExitReal_declaration is called when production real_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitReal_declaration(ctx *Real_declarationContext) {}

// EnterTime_declaration is called when production time_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterTime_declaration(ctx *Time_declarationContext) {}

// ExitTime_declaration is called when production time_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitTime_declaration(ctx *Time_declarationContext) {}

// EnterRealtime_declaration is called when production realtime_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterRealtime_declaration(ctx *Realtime_declarationContext) {}

// ExitRealtime_declaration is called when production realtime_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitRealtime_declaration(ctx *Realtime_declarationContext) {}

// EnterEvent_declaration is called when production event_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_declaration(ctx *Event_declarationContext) {}

// ExitEvent_declaration is called when production event_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_declaration(ctx *Event_declarationContext) {}

// EnterGenvar_declaration is called when production genvar_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterGenvar_declaration(ctx *Genvar_declarationContext) {}

// ExitGenvar_declaration is called when production genvar_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitGenvar_declaration(ctx *Genvar_declarationContext) {}

// EnterUsertype_variable_declaration is called when production usertype_variable_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterUsertype_variable_declaration(ctx *Usertype_variable_declarationContext) {
}

// ExitUsertype_variable_declaration is called when production usertype_variable_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitUsertype_variable_declaration(ctx *Usertype_variable_declarationContext) {
}

// EnterString_declaration is called when production string_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterString_declaration(ctx *String_declarationContext) {}

// ExitString_declaration is called when production string_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitString_declaration(ctx *String_declarationContext) {}

// EnterStruct_declaration is called when production struct_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterStruct_declaration(ctx *Struct_declarationContext) {}

// ExitStruct_declaration is called when production struct_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitStruct_declaration(ctx *Struct_declarationContext) {}

// EnterEnum_declaration is called when production enum_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterEnum_declaration(ctx *Enum_declarationContext) {}

// ExitEnum_declaration is called when production enum_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitEnum_declaration(ctx *Enum_declarationContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterFunction_type is called when production function_type is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_type(ctx *Function_typeContext) {}

// ExitFunction_type is called when production function_type is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_type(ctx *Function_typeContext) {}

// EnterFunction_identifier is called when production function_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_identifier(ctx *Function_identifierContext) {}

// ExitFunction_identifier is called when production function_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_identifier(ctx *Function_identifierContext) {}

// EnterFunction_interface is called when production function_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_interface(ctx *Function_interfaceContext) {}

// ExitFunction_interface is called when production function_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_interface(ctx *Function_interfaceContext) {}

// EnterFunction_item_declaration_star is called when production function_item_declaration_star is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_item_declaration_star(ctx *Function_item_declaration_starContext) {
}

// ExitFunction_item_declaration_star is called when production function_item_declaration_star is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_item_declaration_star(ctx *Function_item_declaration_starContext) {
}

// EnterFunction_item_declaration_semicolon is called when production function_item_declaration_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_item_declaration_semicolon(ctx *Function_item_declaration_semicolonContext) {
}

// ExitFunction_item_declaration_semicolon is called when production function_item_declaration_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_item_declaration_semicolon(ctx *Function_item_declaration_semicolonContext) {
}

// EnterFunction_item_declaration is called when production function_item_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_item_declaration(ctx *Function_item_declarationContext) {
}

// ExitFunction_item_declaration is called when production function_item_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_item_declaration(ctx *Function_item_declarationContext) {
}

// EnterFunction_statement is called when production function_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_statement(ctx *Function_statementContext) {}

// ExitFunction_statement is called when production function_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_statement(ctx *Function_statementContext) {}

// EnterColon_function_identifier is called when production colon_function_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterColon_function_identifier(ctx *Colon_function_identifierContext) {
}

// ExitColon_function_identifier is called when production colon_function_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitColon_function_identifier(ctx *Colon_function_identifierContext) {
}

// EnterTask_declaration is called when production task_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_declaration(ctx *Task_declarationContext) {}

// ExitTask_declaration is called when production task_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_declaration(ctx *Task_declarationContext) {}

// EnterTask_identifier is called when production task_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_identifier(ctx *Task_identifierContext) {}

// ExitTask_identifier is called when production task_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_identifier(ctx *Task_identifierContext) {}

// EnterTask_interface is called when production task_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_interface(ctx *Task_interfaceContext) {}

// ExitTask_interface is called when production task_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_interface(ctx *Task_interfaceContext) {}

// EnterTask_item_declaration_semicolon is called when production task_item_declaration_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_item_declaration_semicolon(ctx *Task_item_declaration_semicolonContext) {
}

// ExitTask_item_declaration_semicolon is called when production task_item_declaration_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_item_declaration_semicolon(ctx *Task_item_declaration_semicolonContext) {
}

// EnterTask_item_declaration is called when production task_item_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_item_declaration(ctx *Task_item_declarationContext) {}

// ExitTask_item_declaration is called when production task_item_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_item_declaration(ctx *Task_item_declarationContext) {}

// EnterTask_item_declaration_star is called when production task_item_declaration_star is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_item_declaration_star(ctx *Task_item_declaration_starContext) {
}

// ExitTask_item_declaration_star is called when production task_item_declaration_star is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_item_declaration_star(ctx *Task_item_declaration_starContext) {
}

// EnterTask_statement is called when production task_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_statement(ctx *Task_statementContext) {}

// ExitTask_statement is called when production task_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_statement(ctx *Task_statementContext) {}

// EnterStruct_item_semicolon is called when production struct_item_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterStruct_item_semicolon(ctx *Struct_item_semicolonContext) {}

// ExitStruct_item_semicolon is called when production struct_item_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitStruct_item_semicolon(ctx *Struct_item_semicolonContext) {}

// EnterStruct_item_star is called when production struct_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterStruct_item_star(ctx *Struct_item_starContext) {}

// ExitStruct_item_star is called when production struct_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitStruct_item_star(ctx *Struct_item_starContext) {}

// EnterStruct_item is called when production struct_item is entered.
func (s *BaseSysVerilogHDLListener) EnterStruct_item(ctx *Struct_itemContext) {}

// ExitStruct_item is called when production struct_item is exited.
func (s *BaseSysVerilogHDLListener) ExitStruct_item(ctx *Struct_itemContext) {}

// EnterStruct_type is called when production struct_type is entered.
func (s *BaseSysVerilogHDLListener) EnterStruct_type(ctx *Struct_typeContext) {}

// ExitStruct_type is called when production struct_type is exited.
func (s *BaseSysVerilogHDLListener) ExitStruct_type(ctx *Struct_typeContext) {}

// EnterEnum_type is called when production enum_type is entered.
func (s *BaseSysVerilogHDLListener) EnterEnum_type(ctx *Enum_typeContext) {}

// ExitEnum_type is called when production enum_type is exited.
func (s *BaseSysVerilogHDLListener) ExitEnum_type(ctx *Enum_typeContext) {}

// EnterList_of_enum_items is called when production list_of_enum_items is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_enum_items(ctx *List_of_enum_itemsContext) {}

// ExitList_of_enum_items is called when production list_of_enum_items is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_enum_items(ctx *List_of_enum_itemsContext) {}

// EnterEnum_item is called when production enum_item is entered.
func (s *BaseSysVerilogHDLListener) EnterEnum_item(ctx *Enum_itemContext) {}

// ExitEnum_item is called when production enum_item is exited.
func (s *BaseSysVerilogHDLListener) ExitEnum_item(ctx *Enum_itemContext) {}

// EnterEnum_identifier is called when production enum_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterEnum_identifier(ctx *Enum_identifierContext) {}

// ExitEnum_identifier is called when production enum_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitEnum_identifier(ctx *Enum_identifierContext) {}

// EnterComma_enum_item_star is called when production comma_enum_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_enum_item_star(ctx *Comma_enum_item_starContext) {}

// ExitComma_enum_item_star is called when production comma_enum_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_enum_item_star(ctx *Comma_enum_item_starContext) {}

// EnterComma_enum_item is called when production comma_enum_item is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_enum_item(ctx *Comma_enum_itemContext) {}

// ExitComma_enum_item is called when production comma_enum_item is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_enum_item(ctx *Comma_enum_itemContext) {}

// EnterEnumerated_type is called when production enumerated_type is entered.
func (s *BaseSysVerilogHDLListener) EnterEnumerated_type(ctx *Enumerated_typeContext) {}

// ExitEnumerated_type is called when production enumerated_type is exited.
func (s *BaseSysVerilogHDLListener) ExitEnumerated_type(ctx *Enumerated_typeContext) {}

// EnterModule_instantiation is called when production module_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_instantiation(ctx *Module_instantiationContext) {}

// ExitModule_instantiation is called when production module_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_instantiation(ctx *Module_instantiationContext) {}

// EnterParameter_interface_assignments is called when production parameter_interface_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterParameter_interface_assignments(ctx *Parameter_interface_assignmentsContext) {
}

// ExitParameter_interface_assignments is called when production parameter_interface_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitParameter_interface_assignments(ctx *Parameter_interface_assignmentsContext) {
}

// EnterList_of_interface_assignments is called when production list_of_interface_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_interface_assignments(ctx *List_of_interface_assignmentsContext) {
}

// ExitList_of_interface_assignments is called when production list_of_interface_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_interface_assignments(ctx *List_of_interface_assignmentsContext) {
}

// EnterList_of_ordered_interface_assignments is called when production list_of_ordered_interface_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_ordered_interface_assignments(ctx *List_of_ordered_interface_assignmentsContext) {
}

// ExitList_of_ordered_interface_assignments is called when production list_of_ordered_interface_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_ordered_interface_assignments(ctx *List_of_ordered_interface_assignmentsContext) {
}

// EnterComma_ordered_interface_assignment_star is called when production comma_ordered_interface_assignment_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_ordered_interface_assignment_star(ctx *Comma_ordered_interface_assignment_starContext) {
}

// ExitComma_ordered_interface_assignment_star is called when production comma_ordered_interface_assignment_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_ordered_interface_assignment_star(ctx *Comma_ordered_interface_assignment_starContext) {
}

// EnterComma_ordered_interface_assignment is called when production comma_ordered_interface_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_ordered_interface_assignment(ctx *Comma_ordered_interface_assignmentContext) {
}

// ExitComma_ordered_interface_assignment is called when production comma_ordered_interface_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_ordered_interface_assignment(ctx *Comma_ordered_interface_assignmentContext) {
}

// EnterOrdered_interface_assignment is called when production ordered_interface_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterOrdered_interface_assignment(ctx *Ordered_interface_assignmentContext) {
}

// ExitOrdered_interface_assignment is called when production ordered_interface_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitOrdered_interface_assignment(ctx *Ordered_interface_assignmentContext) {
}

// EnterList_of_named_interface_assignments is called when production list_of_named_interface_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_named_interface_assignments(ctx *List_of_named_interface_assignmentsContext) {
}

// ExitList_of_named_interface_assignments is called when production list_of_named_interface_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_named_interface_assignments(ctx *List_of_named_interface_assignmentsContext) {
}

// EnterComma_named_interface_assignment_star is called when production comma_named_interface_assignment_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_named_interface_assignment_star(ctx *Comma_named_interface_assignment_starContext) {
}

// ExitComma_named_interface_assignment_star is called when production comma_named_interface_assignment_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_named_interface_assignment_star(ctx *Comma_named_interface_assignment_starContext) {
}

// EnterComma_named_interface_assignment is called when production comma_named_interface_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_named_interface_assignment(ctx *Comma_named_interface_assignmentContext) {
}

// ExitComma_named_interface_assignment is called when production comma_named_interface_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_named_interface_assignment(ctx *Comma_named_interface_assignmentContext) {
}

// EnterNamed_interface_assignment is called when production named_interface_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterNamed_interface_assignment(ctx *Named_interface_assignmentContext) {
}

// ExitNamed_interface_assignment is called when production named_interface_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitNamed_interface_assignment(ctx *Named_interface_assignmentContext) {
}

// EnterList_of_module_instances is called when production list_of_module_instances is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_module_instances(ctx *List_of_module_instancesContext) {
}

// ExitList_of_module_instances is called when production list_of_module_instances is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_module_instances(ctx *List_of_module_instancesContext) {
}

// EnterComma_module_instance_star is called when production comma_module_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_module_instance_star(ctx *Comma_module_instance_starContext) {
}

// ExitComma_module_instance_star is called when production comma_module_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_module_instance_star(ctx *Comma_module_instance_starContext) {
}

// EnterComma_module_instance is called when production comma_module_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_module_instance(ctx *Comma_module_instanceContext) {}

// ExitComma_module_instance is called when production comma_module_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_module_instance(ctx *Comma_module_instanceContext) {}

// EnterModule_instance is called when production module_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_instance(ctx *Module_instanceContext) {}

// ExitModule_instance is called when production module_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_instance(ctx *Module_instanceContext) {}

// EnterModule_instance_identifier is called when production module_instance_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterModule_instance_identifier(ctx *Module_instance_identifierContext) {
}

// ExitModule_instance_identifier is called when production module_instance_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitModule_instance_identifier(ctx *Module_instance_identifierContext) {
}

// EnterArrayed_identifier is called when production arrayed_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterArrayed_identifier(ctx *Arrayed_identifierContext) {}

// ExitArrayed_identifier is called when production arrayed_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitArrayed_identifier(ctx *Arrayed_identifierContext) {}

// EnterSimple_arrayed_identifier is called when production simple_arrayed_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterSimple_arrayed_identifier(ctx *Simple_arrayed_identifierContext) {
}

// ExitSimple_arrayed_identifier is called when production simple_arrayed_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitSimple_arrayed_identifier(ctx *Simple_arrayed_identifierContext) {
}

// EnterEscaped_arrayed_identifier is called when production escaped_arrayed_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterEscaped_arrayed_identifier(ctx *Escaped_arrayed_identifierContext) {
}

// ExitEscaped_arrayed_identifier is called when production escaped_arrayed_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitEscaped_arrayed_identifier(ctx *Escaped_arrayed_identifierContext) {
}

// EnterPort_interface_assignments is called when production port_interface_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterPort_interface_assignments(ctx *Port_interface_assignmentsContext) {
}

// ExitPort_interface_assignments is called when production port_interface_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitPort_interface_assignments(ctx *Port_interface_assignmentsContext) {
}

// EnterDelay is called when production delay is entered.
func (s *BaseSysVerilogHDLListener) EnterDelay(ctx *DelayContext) {}

// ExitDelay is called when production delay is exited.
func (s *BaseSysVerilogHDLListener) ExitDelay(ctx *DelayContext) {}

// EnterList_of_delay_values is called when production list_of_delay_values is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_delay_values(ctx *List_of_delay_valuesContext) {}

// ExitList_of_delay_values is called when production list_of_delay_values is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_delay_values(ctx *List_of_delay_valuesContext) {}

// EnterComma_delay_value_star is called when production comma_delay_value_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_delay_value_star(ctx *Comma_delay_value_starContext) {}

// ExitComma_delay_value_star is called when production comma_delay_value_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_delay_value_star(ctx *Comma_delay_value_starContext) {}

// EnterComma_delay_value is called when production comma_delay_value is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_delay_value(ctx *Comma_delay_valueContext) {}

// ExitComma_delay_value is called when production comma_delay_value is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_delay_value(ctx *Comma_delay_valueContext) {}

// EnterDelay_value is called when production delay_value is entered.
func (s *BaseSysVerilogHDLListener) EnterDelay_value(ctx *Delay_valueContext) {}

// ExitDelay_value is called when production delay_value is exited.
func (s *BaseSysVerilogHDLListener) ExitDelay_value(ctx *Delay_valueContext) {}

// EnterPulldown_strength is called when production pulldown_strength is entered.
func (s *BaseSysVerilogHDLListener) EnterPulldown_strength(ctx *Pulldown_strengthContext) {}

// ExitPulldown_strength is called when production pulldown_strength is exited.
func (s *BaseSysVerilogHDLListener) ExitPulldown_strength(ctx *Pulldown_strengthContext) {}

// EnterPullup_strength is called when production pullup_strength is entered.
func (s *BaseSysVerilogHDLListener) EnterPullup_strength(ctx *Pullup_strengthContext) {}

// ExitPullup_strength is called when production pullup_strength is exited.
func (s *BaseSysVerilogHDLListener) ExitPullup_strength(ctx *Pullup_strengthContext) {}

// EnterGate_instance_identifier is called when production gate_instance_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterGate_instance_identifier(ctx *Gate_instance_identifierContext) {
}

// ExitGate_instance_identifier is called when production gate_instance_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitGate_instance_identifier(ctx *Gate_instance_identifierContext) {
}

// EnterGate_instantiation is called when production gate_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterGate_instantiation(ctx *Gate_instantiationContext) {}

// ExitGate_instantiation is called when production gate_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitGate_instantiation(ctx *Gate_instantiationContext) {}

// EnterEnable_gatetype is called when production enable_gatetype is entered.
func (s *BaseSysVerilogHDLListener) EnterEnable_gatetype(ctx *Enable_gatetypeContext) {}

// ExitEnable_gatetype is called when production enable_gatetype is exited.
func (s *BaseSysVerilogHDLListener) ExitEnable_gatetype(ctx *Enable_gatetypeContext) {}

// EnterMos_switchtype is called when production mos_switchtype is entered.
func (s *BaseSysVerilogHDLListener) EnterMos_switchtype(ctx *Mos_switchtypeContext) {}

// ExitMos_switchtype is called when production mos_switchtype is exited.
func (s *BaseSysVerilogHDLListener) ExitMos_switchtype(ctx *Mos_switchtypeContext) {}

// EnterCmos_switchtype is called when production cmos_switchtype is entered.
func (s *BaseSysVerilogHDLListener) EnterCmos_switchtype(ctx *Cmos_switchtypeContext) {}

// ExitCmos_switchtype is called when production cmos_switchtype is exited.
func (s *BaseSysVerilogHDLListener) ExitCmos_switchtype(ctx *Cmos_switchtypeContext) {}

// EnterN_output_gatetype is called when production n_output_gatetype is entered.
func (s *BaseSysVerilogHDLListener) EnterN_output_gatetype(ctx *N_output_gatetypeContext) {}

// ExitN_output_gatetype is called when production n_output_gatetype is exited.
func (s *BaseSysVerilogHDLListener) ExitN_output_gatetype(ctx *N_output_gatetypeContext) {}

// EnterN_input_gatetype is called when production n_input_gatetype is entered.
func (s *BaseSysVerilogHDLListener) EnterN_input_gatetype(ctx *N_input_gatetypeContext) {}

// ExitN_input_gatetype is called when production n_input_gatetype is exited.
func (s *BaseSysVerilogHDLListener) ExitN_input_gatetype(ctx *N_input_gatetypeContext) {}

// EnterPass_switchtype is called when production pass_switchtype is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_switchtype(ctx *Pass_switchtypeContext) {}

// ExitPass_switchtype is called when production pass_switchtype is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_switchtype(ctx *Pass_switchtypeContext) {}

// EnterPass_enable_switchtype is called when production pass_enable_switchtype is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_enable_switchtype(ctx *Pass_enable_switchtypeContext) {}

// ExitPass_enable_switchtype is called when production pass_enable_switchtype is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_enable_switchtype(ctx *Pass_enable_switchtypeContext) {}

// EnterPulldown_instantiation is called when production pulldown_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterPulldown_instantiation(ctx *Pulldown_instantiationContext) {}

// ExitPulldown_instantiation is called when production pulldown_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitPulldown_instantiation(ctx *Pulldown_instantiationContext) {}

// EnterPullup_instantiation is called when production pullup_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterPullup_instantiation(ctx *Pullup_instantiationContext) {}

// ExitPullup_instantiation is called when production pullup_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitPullup_instantiation(ctx *Pullup_instantiationContext) {}

// EnterEnable_instantiation is called when production enable_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterEnable_instantiation(ctx *Enable_instantiationContext) {}

// ExitEnable_instantiation is called when production enable_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitEnable_instantiation(ctx *Enable_instantiationContext) {}

// EnterMos_instantiation is called when production mos_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterMos_instantiation(ctx *Mos_instantiationContext) {}

// ExitMos_instantiation is called when production mos_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitMos_instantiation(ctx *Mos_instantiationContext) {}

// EnterCmos_instantiation is called when production cmos_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterCmos_instantiation(ctx *Cmos_instantiationContext) {}

// ExitCmos_instantiation is called when production cmos_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitCmos_instantiation(ctx *Cmos_instantiationContext) {}

// EnterN_output_instantiation is called when production n_output_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterN_output_instantiation(ctx *N_output_instantiationContext) {}

// ExitN_output_instantiation is called when production n_output_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitN_output_instantiation(ctx *N_output_instantiationContext) {}

// EnterN_input_instantiation is called when production n_input_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterN_input_instantiation(ctx *N_input_instantiationContext) {}

// ExitN_input_instantiation is called when production n_input_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitN_input_instantiation(ctx *N_input_instantiationContext) {}

// EnterPass_instantiation is called when production pass_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_instantiation(ctx *Pass_instantiationContext) {}

// ExitPass_instantiation is called when production pass_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_instantiation(ctx *Pass_instantiationContext) {}

// EnterPass_enable_instantiation is called when production pass_enable_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_enable_instantiation(ctx *Pass_enable_instantiationContext) {
}

// ExitPass_enable_instantiation is called when production pass_enable_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_enable_instantiation(ctx *Pass_enable_instantiationContext) {
}

// EnterList_of_pull_gate_instance is called when production list_of_pull_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_pull_gate_instance(ctx *List_of_pull_gate_instanceContext) {
}

// ExitList_of_pull_gate_instance is called when production list_of_pull_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_pull_gate_instance(ctx *List_of_pull_gate_instanceContext) {
}

// EnterList_of_enable_gate_instance is called when production list_of_enable_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_enable_gate_instance(ctx *List_of_enable_gate_instanceContext) {
}

// ExitList_of_enable_gate_instance is called when production list_of_enable_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_enable_gate_instance(ctx *List_of_enable_gate_instanceContext) {
}

// EnterList_of_mos_switch_instance is called when production list_of_mos_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_mos_switch_instance(ctx *List_of_mos_switch_instanceContext) {
}

// ExitList_of_mos_switch_instance is called when production list_of_mos_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_mos_switch_instance(ctx *List_of_mos_switch_instanceContext) {
}

// EnterList_of_cmos_switch_instance is called when production list_of_cmos_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_cmos_switch_instance(ctx *List_of_cmos_switch_instanceContext) {
}

// ExitList_of_cmos_switch_instance is called when production list_of_cmos_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_cmos_switch_instance(ctx *List_of_cmos_switch_instanceContext) {
}

// EnterList_of_n_input_gate_instance is called when production list_of_n_input_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_n_input_gate_instance(ctx *List_of_n_input_gate_instanceContext) {
}

// ExitList_of_n_input_gate_instance is called when production list_of_n_input_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_n_input_gate_instance(ctx *List_of_n_input_gate_instanceContext) {
}

// EnterList_of_n_output_gate_instance is called when production list_of_n_output_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_n_output_gate_instance(ctx *List_of_n_output_gate_instanceContext) {
}

// ExitList_of_n_output_gate_instance is called when production list_of_n_output_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_n_output_gate_instance(ctx *List_of_n_output_gate_instanceContext) {
}

// EnterList_of_pass_switch_instance is called when production list_of_pass_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_pass_switch_instance(ctx *List_of_pass_switch_instanceContext) {
}

// ExitList_of_pass_switch_instance is called when production list_of_pass_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_pass_switch_instance(ctx *List_of_pass_switch_instanceContext) {
}

// EnterList_of_pass_enable_switch_instance is called when production list_of_pass_enable_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_pass_enable_switch_instance(ctx *List_of_pass_enable_switch_instanceContext) {
}

// ExitList_of_pass_enable_switch_instance is called when production list_of_pass_enable_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_pass_enable_switch_instance(ctx *List_of_pass_enable_switch_instanceContext) {
}

// EnterComma_pull_gate_instance_star is called when production comma_pull_gate_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_pull_gate_instance_star(ctx *Comma_pull_gate_instance_starContext) {
}

// ExitComma_pull_gate_instance_star is called when production comma_pull_gate_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_pull_gate_instance_star(ctx *Comma_pull_gate_instance_starContext) {
}

// EnterComma_enable_gate_instance_star is called when production comma_enable_gate_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_enable_gate_instance_star(ctx *Comma_enable_gate_instance_starContext) {
}

// ExitComma_enable_gate_instance_star is called when production comma_enable_gate_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_enable_gate_instance_star(ctx *Comma_enable_gate_instance_starContext) {
}

// EnterComma_mos_switch_instance_star is called when production comma_mos_switch_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_mos_switch_instance_star(ctx *Comma_mos_switch_instance_starContext) {
}

// ExitComma_mos_switch_instance_star is called when production comma_mos_switch_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_mos_switch_instance_star(ctx *Comma_mos_switch_instance_starContext) {
}

// EnterComma_cmos_switch_instance_star is called when production comma_cmos_switch_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_cmos_switch_instance_star(ctx *Comma_cmos_switch_instance_starContext) {
}

// ExitComma_cmos_switch_instance_star is called when production comma_cmos_switch_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_cmos_switch_instance_star(ctx *Comma_cmos_switch_instance_starContext) {
}

// EnterComma_n_input_gate_instance_star is called when production comma_n_input_gate_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_n_input_gate_instance_star(ctx *Comma_n_input_gate_instance_starContext) {
}

// ExitComma_n_input_gate_instance_star is called when production comma_n_input_gate_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_n_input_gate_instance_star(ctx *Comma_n_input_gate_instance_starContext) {
}

// EnterComma_n_output_gate_instance_star is called when production comma_n_output_gate_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_n_output_gate_instance_star(ctx *Comma_n_output_gate_instance_starContext) {
}

// ExitComma_n_output_gate_instance_star is called when production comma_n_output_gate_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_n_output_gate_instance_star(ctx *Comma_n_output_gate_instance_starContext) {
}

// EnterComma_pass_switch_instance_star is called when production comma_pass_switch_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_pass_switch_instance_star(ctx *Comma_pass_switch_instance_starContext) {
}

// ExitComma_pass_switch_instance_star is called when production comma_pass_switch_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_pass_switch_instance_star(ctx *Comma_pass_switch_instance_starContext) {
}

// EnterComma_pass_enable_switch_instance_star is called when production comma_pass_enable_switch_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_pass_enable_switch_instance_star(ctx *Comma_pass_enable_switch_instance_starContext) {
}

// ExitComma_pass_enable_switch_instance_star is called when production comma_pass_enable_switch_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_pass_enable_switch_instance_star(ctx *Comma_pass_enable_switch_instance_starContext) {
}

// EnterComma_pull_gate_instance is called when production comma_pull_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_pull_gate_instance(ctx *Comma_pull_gate_instanceContext) {
}

// ExitComma_pull_gate_instance is called when production comma_pull_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_pull_gate_instance(ctx *Comma_pull_gate_instanceContext) {
}

// EnterComma_enable_gate_instance is called when production comma_enable_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_enable_gate_instance(ctx *Comma_enable_gate_instanceContext) {
}

// ExitComma_enable_gate_instance is called when production comma_enable_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_enable_gate_instance(ctx *Comma_enable_gate_instanceContext) {
}

// EnterComma_mos_switch_instance is called when production comma_mos_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_mos_switch_instance(ctx *Comma_mos_switch_instanceContext) {
}

// ExitComma_mos_switch_instance is called when production comma_mos_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_mos_switch_instance(ctx *Comma_mos_switch_instanceContext) {
}

// EnterComma_cmos_switch_instance is called when production comma_cmos_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_cmos_switch_instance(ctx *Comma_cmos_switch_instanceContext) {
}

// ExitComma_cmos_switch_instance is called when production comma_cmos_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_cmos_switch_instance(ctx *Comma_cmos_switch_instanceContext) {
}

// EnterComma_n_input_gate_instance is called when production comma_n_input_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_n_input_gate_instance(ctx *Comma_n_input_gate_instanceContext) {
}

// ExitComma_n_input_gate_instance is called when production comma_n_input_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_n_input_gate_instance(ctx *Comma_n_input_gate_instanceContext) {
}

// EnterComma_n_output_gate_instance is called when production comma_n_output_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_n_output_gate_instance(ctx *Comma_n_output_gate_instanceContext) {
}

// ExitComma_n_output_gate_instance is called when production comma_n_output_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_n_output_gate_instance(ctx *Comma_n_output_gate_instanceContext) {
}

// EnterComma_pass_switch_instance is called when production comma_pass_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_pass_switch_instance(ctx *Comma_pass_switch_instanceContext) {
}

// ExitComma_pass_switch_instance is called when production comma_pass_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_pass_switch_instance(ctx *Comma_pass_switch_instanceContext) {
}

// EnterComma_pass_enable_switch_instance is called when production comma_pass_enable_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_pass_enable_switch_instance(ctx *Comma_pass_enable_switch_instanceContext) {
}

// ExitComma_pass_enable_switch_instance is called when production comma_pass_enable_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_pass_enable_switch_instance(ctx *Comma_pass_enable_switch_instanceContext) {
}

// EnterPull_gate_instance is called when production pull_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterPull_gate_instance(ctx *Pull_gate_instanceContext) {}

// ExitPull_gate_instance is called when production pull_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitPull_gate_instance(ctx *Pull_gate_instanceContext) {}

// EnterEnable_gate_instance is called when production enable_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterEnable_gate_instance(ctx *Enable_gate_instanceContext) {}

// ExitEnable_gate_instance is called when production enable_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitEnable_gate_instance(ctx *Enable_gate_instanceContext) {}

// EnterMos_switch_instance is called when production mos_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterMos_switch_instance(ctx *Mos_switch_instanceContext) {}

// ExitMos_switch_instance is called when production mos_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitMos_switch_instance(ctx *Mos_switch_instanceContext) {}

// EnterCmos_switch_instance is called when production cmos_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterCmos_switch_instance(ctx *Cmos_switch_instanceContext) {}

// ExitCmos_switch_instance is called when production cmos_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitCmos_switch_instance(ctx *Cmos_switch_instanceContext) {}

// EnterN_input_gate_instance is called when production n_input_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterN_input_gate_instance(ctx *N_input_gate_instanceContext) {}

// ExitN_input_gate_instance is called when production n_input_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitN_input_gate_instance(ctx *N_input_gate_instanceContext) {}

// EnterN_output_gate_instance is called when production n_output_gate_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterN_output_gate_instance(ctx *N_output_gate_instanceContext) {}

// ExitN_output_gate_instance is called when production n_output_gate_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitN_output_gate_instance(ctx *N_output_gate_instanceContext) {}

// EnterPass_switch_instance is called when production pass_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_switch_instance(ctx *Pass_switch_instanceContext) {}

// ExitPass_switch_instance is called when production pass_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_switch_instance(ctx *Pass_switch_instanceContext) {}

// EnterPass_enable_switch_instance is called when production pass_enable_switch_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_enable_switch_instance(ctx *Pass_enable_switch_instanceContext) {
}

// ExitPass_enable_switch_instance is called when production pass_enable_switch_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_enable_switch_instance(ctx *Pass_enable_switch_instanceContext) {
}

// EnterPull_gate_interface is called when production pull_gate_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterPull_gate_interface(ctx *Pull_gate_interfaceContext) {}

// ExitPull_gate_interface is called when production pull_gate_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitPull_gate_interface(ctx *Pull_gate_interfaceContext) {}

// EnterEnable_gate_interface is called when production enable_gate_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterEnable_gate_interface(ctx *Enable_gate_interfaceContext) {}

// ExitEnable_gate_interface is called when production enable_gate_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitEnable_gate_interface(ctx *Enable_gate_interfaceContext) {}

// EnterMos_switch_interface is called when production mos_switch_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterMos_switch_interface(ctx *Mos_switch_interfaceContext) {}

// ExitMos_switch_interface is called when production mos_switch_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitMos_switch_interface(ctx *Mos_switch_interfaceContext) {}

// EnterCmos_switch_interface is called when production cmos_switch_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterCmos_switch_interface(ctx *Cmos_switch_interfaceContext) {}

// ExitCmos_switch_interface is called when production cmos_switch_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitCmos_switch_interface(ctx *Cmos_switch_interfaceContext) {}

// EnterN_input_gate_interface is called when production n_input_gate_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterN_input_gate_interface(ctx *N_input_gate_interfaceContext) {}

// ExitN_input_gate_interface is called when production n_input_gate_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitN_input_gate_interface(ctx *N_input_gate_interfaceContext) {}

// EnterN_output_gate_interface is called when production n_output_gate_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterN_output_gate_interface(ctx *N_output_gate_interfaceContext) {
}

// ExitN_output_gate_interface is called when production n_output_gate_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitN_output_gate_interface(ctx *N_output_gate_interfaceContext) {}

// EnterPass_switch_interface is called when production pass_switch_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_switch_interface(ctx *Pass_switch_interfaceContext) {}

// ExitPass_switch_interface is called when production pass_switch_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_switch_interface(ctx *Pass_switch_interfaceContext) {}

// EnterPass_enable_switch_interface is called when production pass_enable_switch_interface is entered.
func (s *BaseSysVerilogHDLListener) EnterPass_enable_switch_interface(ctx *Pass_enable_switch_interfaceContext) {
}

// ExitPass_enable_switch_interface is called when production pass_enable_switch_interface is exited.
func (s *BaseSysVerilogHDLListener) ExitPass_enable_switch_interface(ctx *Pass_enable_switch_interfaceContext) {
}

// EnterList_of_input_terminals is called when production list_of_input_terminals is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_input_terminals(ctx *List_of_input_terminalsContext) {
}

// ExitList_of_input_terminals is called when production list_of_input_terminals is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_input_terminals(ctx *List_of_input_terminalsContext) {}

// EnterList_of_output_terminals is called when production list_of_output_terminals is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_output_terminals(ctx *List_of_output_terminalsContext) {
}

// ExitList_of_output_terminals is called when production list_of_output_terminals is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_output_terminals(ctx *List_of_output_terminalsContext) {
}

// EnterComma_input_terminal_star is called when production comma_input_terminal_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_input_terminal_star(ctx *Comma_input_terminal_starContext) {
}

// ExitComma_input_terminal_star is called when production comma_input_terminal_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_input_terminal_star(ctx *Comma_input_terminal_starContext) {
}

// EnterComma_output_terminal_star is called when production comma_output_terminal_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_output_terminal_star(ctx *Comma_output_terminal_starContext) {
}

// ExitComma_output_terminal_star is called when production comma_output_terminal_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_output_terminal_star(ctx *Comma_output_terminal_starContext) {
}

// EnterComma_input_terminal is called when production comma_input_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_input_terminal(ctx *Comma_input_terminalContext) {}

// ExitComma_input_terminal is called when production comma_input_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_input_terminal(ctx *Comma_input_terminalContext) {}

// EnterComma_output_terminal is called when production comma_output_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_output_terminal(ctx *Comma_output_terminalContext) {}

// ExitComma_output_terminal is called when production comma_output_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_output_terminal(ctx *Comma_output_terminalContext) {}

// EnterEnable_terminal is called when production enable_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterEnable_terminal(ctx *Enable_terminalContext) {}

// ExitEnable_terminal is called when production enable_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitEnable_terminal(ctx *Enable_terminalContext) {}

// EnterInput_terminal is called when production input_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterInput_terminal(ctx *Input_terminalContext) {}

// ExitInput_terminal is called when production input_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitInput_terminal(ctx *Input_terminalContext) {}

// EnterInout_terminal is called when production inout_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterInout_terminal(ctx *Inout_terminalContext) {}

// ExitInout_terminal is called when production inout_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitInout_terminal(ctx *Inout_terminalContext) {}

// EnterNcontrol_terminal is called when production ncontrol_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterNcontrol_terminal(ctx *Ncontrol_terminalContext) {}

// ExitNcontrol_terminal is called when production ncontrol_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitNcontrol_terminal(ctx *Ncontrol_terminalContext) {}

// EnterOutput_terminal is called when production output_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterOutput_terminal(ctx *Output_terminalContext) {}

// ExitOutput_terminal is called when production output_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitOutput_terminal(ctx *Output_terminalContext) {}

// EnterPcontrol_terminal is called when production pcontrol_terminal is entered.
func (s *BaseSysVerilogHDLListener) EnterPcontrol_terminal(ctx *Pcontrol_terminalContext) {}

// ExitPcontrol_terminal is called when production pcontrol_terminal is exited.
func (s *BaseSysVerilogHDLListener) ExitPcontrol_terminal(ctx *Pcontrol_terminalContext) {}

// EnterStatement_star is called when production statement_star is entered.
func (s *BaseSysVerilogHDLListener) EnterStatement_star(ctx *Statement_starContext) {}

// ExitStatement_star is called when production statement_star is exited.
func (s *BaseSysVerilogHDLListener) ExitStatement_star(ctx *Statement_starContext) {}

// EnterStatement_semicolon is called when production statement_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterStatement_semicolon(ctx *Statement_semicolonContext) {}

// ExitStatement_semicolon is called when production statement_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitStatement_semicolon(ctx *Statement_semicolonContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSysVerilogHDLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSysVerilogHDLListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterFlow_control_statement is called when production flow_control_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterFlow_control_statement(ctx *Flow_control_statementContext) {}

// ExitFlow_control_statement is called when production flow_control_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitFlow_control_statement(ctx *Flow_control_statementContext) {}

// EnterBlock_statement is called when production block_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterBlock_statement(ctx *Block_statementContext) {}

// ExitBlock_statement is called when production block_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitBlock_statement(ctx *Block_statementContext) {}

// EnterTask_call_statement is called when production task_call_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_call_statement(ctx *Task_call_statementContext) {}

// ExitTask_call_statement is called when production task_call_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_call_statement(ctx *Task_call_statementContext) {}

// EnterEvent_statement is called when production event_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_statement(ctx *Event_statementContext) {}

// ExitEvent_statement is called when production event_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_statement(ctx *Event_statementContext) {}

// EnterProcedural_statement is called when production procedural_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterProcedural_statement(ctx *Procedural_statementContext) {}

// ExitProcedural_statement is called when production procedural_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitProcedural_statement(ctx *Procedural_statementContext) {}

// EnterExpression_statement is called when production expression_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterExpression_statement(ctx *Expression_statementContext) {}

// ExitExpression_statement is called when production expression_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitExpression_statement(ctx *Expression_statementContext) {}

// EnterSubroutine_statement is called when production subroutine_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterSubroutine_statement(ctx *Subroutine_statementContext) {}

// ExitSubroutine_statement is called when production subroutine_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitSubroutine_statement(ctx *Subroutine_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterNull_statement is called when production null_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterNull_statement(ctx *Null_statementContext) {}

// ExitNull_statement is called when production null_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitNull_statement(ctx *Null_statementContext) {}

// EnterProcedural_continuous_assignments is called when production procedural_continuous_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterProcedural_continuous_assignments(ctx *Procedural_continuous_assignmentsContext) {
}

// ExitProcedural_continuous_assignments is called when production procedural_continuous_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitProcedural_continuous_assignments(ctx *Procedural_continuous_assignmentsContext) {
}

// EnterAssign_statement is called when production assign_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterAssign_statement(ctx *Assign_statementContext) {}

// ExitAssign_statement is called when production assign_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitAssign_statement(ctx *Assign_statementContext) {}

// EnterDeassign_statement is called when production deassign_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterDeassign_statement(ctx *Deassign_statementContext) {}

// ExitDeassign_statement is called when production deassign_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitDeassign_statement(ctx *Deassign_statementContext) {}

// EnterForce_statement is called when production force_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterForce_statement(ctx *Force_statementContext) {}

// ExitForce_statement is called when production force_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitForce_statement(ctx *Force_statementContext) {}

// EnterRelease_statement is called when production release_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterRelease_statement(ctx *Release_statementContext) {}

// ExitRelease_statement is called when production release_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitRelease_statement(ctx *Release_statementContext) {}

// EnterProcedural_timing_control_statement is called when production procedural_timing_control_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterProcedural_timing_control_statement(ctx *Procedural_timing_control_statementContext) {
}

// ExitProcedural_timing_control_statement is called when production procedural_timing_control_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitProcedural_timing_control_statement(ctx *Procedural_timing_control_statementContext) {
}

// EnterProperty_statement is called when production property_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterProperty_statement(ctx *Property_statementContext) {}

// ExitProperty_statement is called when production property_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitProperty_statement(ctx *Property_statementContext) {}

// EnterDisable_condition_statement is called when production disable_condition_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterDisable_condition_statement(ctx *Disable_condition_statementContext) {
}

// ExitDisable_condition_statement is called when production disable_condition_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitDisable_condition_statement(ctx *Disable_condition_statementContext) {
}

// EnterProperty_expression is called when production property_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterProperty_expression(ctx *Property_expressionContext) {}

// ExitProperty_expression is called when production property_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitProperty_expression(ctx *Property_expressionContext) {}

// EnterProcedural_assertion_statement is called when production procedural_assertion_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterProcedural_assertion_statement(ctx *Procedural_assertion_statementContext) {
}

// ExitProcedural_assertion_statement is called when production procedural_assertion_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitProcedural_assertion_statement(ctx *Procedural_assertion_statementContext) {
}

// EnterAssert_else_statement is called when production assert_else_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterAssert_else_statement(ctx *Assert_else_statementContext) {}

// ExitAssert_else_statement is called when production assert_else_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitAssert_else_statement(ctx *Assert_else_statementContext) {}

// EnterAssert_statement is called when production assert_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterAssert_statement(ctx *Assert_statementContext) {}

// ExitAssert_statement is called when production assert_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitAssert_statement(ctx *Assert_statementContext) {}

// EnterSystem_task_enable is called when production system_task_enable is entered.
func (s *BaseSysVerilogHDLListener) EnterSystem_task_enable(ctx *System_task_enableContext) {}

// ExitSystem_task_enable is called when production system_task_enable is exited.
func (s *BaseSysVerilogHDLListener) ExitSystem_task_enable(ctx *System_task_enableContext) {}

// EnterSystem_task_identifier is called when production system_task_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterSystem_task_identifier(ctx *System_task_identifierContext) {}

// ExitSystem_task_identifier is called when production system_task_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitSystem_task_identifier(ctx *System_task_identifierContext) {}

// EnterTask_interface_assignments is called when production task_interface_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_interface_assignments(ctx *Task_interface_assignmentsContext) {
}

// ExitTask_interface_assignments is called when production task_interface_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_interface_assignments(ctx *Task_interface_assignmentsContext) {
}

// EnterTask_enable is called when production task_enable is entered.
func (s *BaseSysVerilogHDLListener) EnterTask_enable(ctx *Task_enableContext) {}

// ExitTask_enable is called when production task_enable is exited.
func (s *BaseSysVerilogHDLListener) ExitTask_enable(ctx *Task_enableContext) {}

// EnterHierarchical_task_identifier is called when production hierarchical_task_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_task_identifier(ctx *Hierarchical_task_identifierContext) {
}

// ExitHierarchical_task_identifier is called when production hierarchical_task_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_task_identifier(ctx *Hierarchical_task_identifierContext) {
}

// EnterDisable_statement is called when production disable_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterDisable_statement(ctx *Disable_statementContext) {}

// ExitDisable_statement is called when production disable_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitDisable_statement(ctx *Disable_statementContext) {}

// EnterHierarchical_block_identifier is called when production hierarchical_block_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_block_identifier(ctx *Hierarchical_block_identifierContext) {
}

// ExitHierarchical_block_identifier is called when production hierarchical_block_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_block_identifier(ctx *Hierarchical_block_identifierContext) {
}

// EnterVariable_lvalue is called when production variable_lvalue is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_lvalue(ctx *Variable_lvalueContext) {}

// ExitVariable_lvalue is called when production variable_lvalue is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_lvalue(ctx *Variable_lvalueContext) {}

// EnterHierarchical_variable_lvalue is called when production hierarchical_variable_lvalue is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_variable_lvalue(ctx *Hierarchical_variable_lvalueContext) {
}

// ExitHierarchical_variable_lvalue is called when production hierarchical_variable_lvalue is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_variable_lvalue(ctx *Hierarchical_variable_lvalueContext) {
}

// EnterVariable_concatenation is called when production variable_concatenation is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_concatenation(ctx *Variable_concatenationContext) {}

// ExitVariable_concatenation is called when production variable_concatenation is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_concatenation(ctx *Variable_concatenationContext) {}

// EnterVariable_concatenation_value is called when production variable_concatenation_value is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_concatenation_value(ctx *Variable_concatenation_valueContext) {
}

// ExitVariable_concatenation_value is called when production variable_concatenation_value is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_concatenation_value(ctx *Variable_concatenation_valueContext) {
}

// EnterComma_vcv_star is called when production comma_vcv_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_vcv_star(ctx *Comma_vcv_starContext) {}

// ExitComma_vcv_star is called when production comma_vcv_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_vcv_star(ctx *Comma_vcv_starContext) {}

// EnterBlocking_assignment is called when production blocking_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterBlocking_assignment(ctx *Blocking_assignmentContext) {}

// ExitBlocking_assignment is called when production blocking_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitBlocking_assignment(ctx *Blocking_assignmentContext) {}

// EnterNonblocking_assignment is called when production nonblocking_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterNonblocking_assignment(ctx *Nonblocking_assignmentContext) {}

// ExitNonblocking_assignment is called when production nonblocking_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitNonblocking_assignment(ctx *Nonblocking_assignmentContext) {}

// EnterPrefix_assignment is called when production prefix_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterPrefix_assignment(ctx *Prefix_assignmentContext) {}

// ExitPrefix_assignment is called when production prefix_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitPrefix_assignment(ctx *Prefix_assignmentContext) {}

// EnterPostfix_assignment is called when production postfix_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterPostfix_assignment(ctx *Postfix_assignmentContext) {}

// ExitPostfix_assignment is called when production postfix_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitPostfix_assignment(ctx *Postfix_assignmentContext) {}

// EnterOperator_assignment is called when production operator_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterOperator_assignment(ctx *Operator_assignmentContext) {}

// ExitOperator_assignment is called when production operator_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitOperator_assignment(ctx *Operator_assignmentContext) {}

// EnterDeclarative_assignment is called when production declarative_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterDeclarative_assignment(ctx *Declarative_assignmentContext) {}

// ExitDeclarative_assignment is called when production declarative_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitDeclarative_assignment(ctx *Declarative_assignmentContext) {}

// EnterDelay_or_event_control is called when production delay_or_event_control is entered.
func (s *BaseSysVerilogHDLListener) EnterDelay_or_event_control(ctx *Delay_or_event_controlContext) {}

// ExitDelay_or_event_control is called when production delay_or_event_control is exited.
func (s *BaseSysVerilogHDLListener) ExitDelay_or_event_control(ctx *Delay_or_event_controlContext) {}

// EnterDelay_control is called when production delay_control is entered.
func (s *BaseSysVerilogHDLListener) EnterDelay_control(ctx *Delay_controlContext) {}

// ExitDelay_control is called when production delay_control is exited.
func (s *BaseSysVerilogHDLListener) ExitDelay_control(ctx *Delay_controlContext) {}

// EnterEvent_control is called when production event_control is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_control(ctx *Event_controlContext) {}

// ExitEvent_control is called when production event_control is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_control(ctx *Event_controlContext) {}

// EnterEvent_control_identifier is called when production event_control_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_control_identifier(ctx *Event_control_identifierContext) {
}

// ExitEvent_control_identifier is called when production event_control_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_control_identifier(ctx *Event_control_identifierContext) {
}

// EnterEvent_control_expression is called when production event_control_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_control_expression(ctx *Event_control_expressionContext) {
}

// ExitEvent_control_expression is called when production event_control_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_control_expression(ctx *Event_control_expressionContext) {
}

// EnterEvent_expression is called when production event_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_expression(ctx *Event_expressionContext) {}

// ExitEvent_expression is called when production event_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_expression(ctx *Event_expressionContext) {}

// EnterSingle_event_expression is called when production single_event_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterSingle_event_expression(ctx *Single_event_expressionContext) {
}

// ExitSingle_event_expression is called when production single_event_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitSingle_event_expression(ctx *Single_event_expressionContext) {}

// EnterEvent_expression_edgespec is called when production event_expression_edgespec is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_expression_edgespec(ctx *Event_expression_edgespecContext) {
}

// ExitEvent_expression_edgespec is called when production event_expression_edgespec is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_expression_edgespec(ctx *Event_expression_edgespecContext) {
}

// EnterEvent_expression_or is called when production event_expression_or is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_expression_or(ctx *Event_expression_orContext) {}

// ExitEvent_expression_or is called when production event_expression_or is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_expression_or(ctx *Event_expression_orContext) {}

// EnterList_of_event_expression_comma is called when production list_of_event_expression_comma is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_event_expression_comma(ctx *List_of_event_expression_commaContext) {
}

// ExitList_of_event_expression_comma is called when production list_of_event_expression_comma is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_event_expression_comma(ctx *List_of_event_expression_commaContext) {
}

// EnterComma_event_expression_star is called when production comma_event_expression_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_event_expression_star(ctx *Comma_event_expression_starContext) {
}

// ExitComma_event_expression_star is called when production comma_event_expression_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_event_expression_star(ctx *Comma_event_expression_starContext) {
}

// EnterComma_event_expression is called when production comma_event_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_event_expression(ctx *Comma_event_expressionContext) {}

// ExitComma_event_expression is called when production comma_event_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_event_expression(ctx *Comma_event_expressionContext) {}

// EnterList_of_event_expression_or is called when production list_of_event_expression_or is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_event_expression_or(ctx *List_of_event_expression_orContext) {
}

// ExitList_of_event_expression_or is called when production list_of_event_expression_or is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_event_expression_or(ctx *List_of_event_expression_orContext) {
}

// EnterOr_event_expression_star is called when production or_event_expression_star is entered.
func (s *BaseSysVerilogHDLListener) EnterOr_event_expression_star(ctx *Or_event_expression_starContext) {
}

// ExitOr_event_expression_star is called when production or_event_expression_star is exited.
func (s *BaseSysVerilogHDLListener) ExitOr_event_expression_star(ctx *Or_event_expression_starContext) {
}

// EnterOr_event_expression is called when production or_event_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterOr_event_expression(ctx *Or_event_expressionContext) {}

// ExitOr_event_expression is called when production or_event_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitOr_event_expression(ctx *Or_event_expressionContext) {}

// EnterEvent_control_wildcard is called when production event_control_wildcard is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_control_wildcard(ctx *Event_control_wildcardContext) {}

// ExitEvent_control_wildcard is called when production event_control_wildcard is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_control_wildcard(ctx *Event_control_wildcardContext) {}

// EnterRepeat_event_control is called when production repeat_event_control is entered.
func (s *BaseSysVerilogHDLListener) EnterRepeat_event_control(ctx *Repeat_event_controlContext) {}

// ExitRepeat_event_control is called when production repeat_event_control is exited.
func (s *BaseSysVerilogHDLListener) ExitRepeat_event_control(ctx *Repeat_event_controlContext) {}

// EnterEvent_trigger is called when production event_trigger is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_trigger(ctx *Event_triggerContext) {}

// ExitEvent_trigger is called when production event_trigger is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_trigger(ctx *Event_triggerContext) {}

// EnterHierarchical_event_identifier is called when production hierarchical_event_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_event_identifier(ctx *Hierarchical_event_identifierContext) {
}

// ExitHierarchical_event_identifier is called when production hierarchical_event_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_event_identifier(ctx *Hierarchical_event_identifierContext) {
}

// EnterEvent_identifier is called when production event_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterEvent_identifier(ctx *Event_identifierContext) {}

// ExitEvent_identifier is called when production event_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitEvent_identifier(ctx *Event_identifierContext) {}

// EnterWait_statement is called when production wait_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterWait_statement(ctx *Wait_statementContext) {}

// ExitWait_statement is called when production wait_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitWait_statement(ctx *Wait_statementContext) {}

// EnterAttr_generated_instantiation is called when production attr_generated_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_generated_instantiation(ctx *Attr_generated_instantiationContext) {
}

// ExitAttr_generated_instantiation is called when production attr_generated_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_generated_instantiation(ctx *Attr_generated_instantiationContext) {
}

// EnterGenerated_instantiation is called when production generated_instantiation is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerated_instantiation(ctx *Generated_instantiationContext) {
}

// ExitGenerated_instantiation is called when production generated_instantiation is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerated_instantiation(ctx *Generated_instantiationContext) {}

// EnterGenerate_item_star is called when production generate_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_item_star(ctx *Generate_item_starContext) {}

// ExitGenerate_item_star is called when production generate_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_item_star(ctx *Generate_item_starContext) {}

// EnterGenerate_item is called when production generate_item is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_item(ctx *Generate_itemContext) {}

// ExitGenerate_item is called when production generate_item is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_item(ctx *Generate_itemContext) {}

// EnterGenerate_block is called when production generate_block is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_block(ctx *Generate_blockContext) {}

// ExitGenerate_block is called when production generate_block is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_block(ctx *Generate_blockContext) {}

// EnterGenerate_colon_block_identifier0 is called when production generate_colon_block_identifier0 is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_colon_block_identifier0(ctx *Generate_colon_block_identifier0Context) {
}

// ExitGenerate_colon_block_identifier0 is called when production generate_colon_block_identifier0 is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_colon_block_identifier0(ctx *Generate_colon_block_identifier0Context) {
}

// EnterGenerate_colon_block_identifier1 is called when production generate_colon_block_identifier1 is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_colon_block_identifier1(ctx *Generate_colon_block_identifier1Context) {
}

// ExitGenerate_colon_block_identifier1 is called when production generate_colon_block_identifier1 is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_colon_block_identifier1(ctx *Generate_colon_block_identifier1Context) {
}

// EnterGenerate_colon_block_identifier is called when production generate_colon_block_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_colon_block_identifier(ctx *Generate_colon_block_identifierContext) {
}

// ExitGenerate_colon_block_identifier is called when production generate_colon_block_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_colon_block_identifier(ctx *Generate_colon_block_identifierContext) {
}

// EnterGenerate_block_identifier is called when production generate_block_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_block_identifier(ctx *Generate_block_identifierContext) {
}

// ExitGenerate_block_identifier is called when production generate_block_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_block_identifier(ctx *Generate_block_identifierContext) {
}

// EnterGenerate_conditional_statement is called when production generate_conditional_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_conditional_statement(ctx *Generate_conditional_statementContext) {
}

// ExitGenerate_conditional_statement is called when production generate_conditional_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_conditional_statement(ctx *Generate_conditional_statementContext) {
}

// EnterGenerate_if_statement is called when production generate_if_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_if_statement(ctx *Generate_if_statementContext) {}

// ExitGenerate_if_statement is called when production generate_if_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_if_statement(ctx *Generate_if_statementContext) {}

// EnterGenerate_else_statement is called when production generate_else_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_else_statement(ctx *Generate_else_statementContext) {
}

// ExitGenerate_else_statement is called when production generate_else_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_else_statement(ctx *Generate_else_statementContext) {}

// EnterGenerate_loop_statement is called when production generate_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_loop_statement(ctx *Generate_loop_statementContext) {
}

// ExitGenerate_loop_statement is called when production generate_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_loop_statement(ctx *Generate_loop_statementContext) {}

// EnterGenerate_forever_loop_statement is called when production generate_forever_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_forever_loop_statement(ctx *Generate_forever_loop_statementContext) {
}

// ExitGenerate_forever_loop_statement is called when production generate_forever_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_forever_loop_statement(ctx *Generate_forever_loop_statementContext) {
}

// EnterGenerate_repeat_loop_statement is called when production generate_repeat_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_repeat_loop_statement(ctx *Generate_repeat_loop_statementContext) {
}

// ExitGenerate_repeat_loop_statement is called when production generate_repeat_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_repeat_loop_statement(ctx *Generate_repeat_loop_statementContext) {
}

// EnterGenerate_while_loop_statement is called when production generate_while_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_while_loop_statement(ctx *Generate_while_loop_statementContext) {
}

// ExitGenerate_while_loop_statement is called when production generate_while_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_while_loop_statement(ctx *Generate_while_loop_statementContext) {
}

// EnterGenerate_do_loop_statement is called when production generate_do_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_do_loop_statement(ctx *Generate_do_loop_statementContext) {
}

// ExitGenerate_do_loop_statement is called when production generate_do_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_do_loop_statement(ctx *Generate_do_loop_statementContext) {
}

// EnterGenerate_for_loop_statement is called when production generate_for_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_for_loop_statement(ctx *Generate_for_loop_statementContext) {
}

// ExitGenerate_for_loop_statement is called when production generate_for_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_for_loop_statement(ctx *Generate_for_loop_statementContext) {
}

// EnterGenerate_case_statement is called when production generate_case_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_case_statement(ctx *Generate_case_statementContext) {
}

// ExitGenerate_case_statement is called when production generate_case_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_case_statement(ctx *Generate_case_statementContext) {}

// EnterGenerate_case_item_star is called when production generate_case_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_case_item_star(ctx *Generate_case_item_starContext) {
}

// ExitGenerate_case_item_star is called when production generate_case_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_case_item_star(ctx *Generate_case_item_starContext) {}

// EnterGenerate_case_item is called when production generate_case_item is entered.
func (s *BaseSysVerilogHDLListener) EnterGenerate_case_item(ctx *Generate_case_itemContext) {}

// ExitGenerate_case_item is called when production generate_case_item is exited.
func (s *BaseSysVerilogHDLListener) ExitGenerate_case_item(ctx *Generate_case_itemContext) {}

// EnterConditional_statement is called when production conditional_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterConditional_statement(ctx *Conditional_statementContext) {}

// ExitConditional_statement is called when production conditional_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitConditional_statement(ctx *Conditional_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterElse_statement is called when production else_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterElse_statement(ctx *Else_statementContext) {}

// ExitElse_statement is called when production else_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitElse_statement(ctx *Else_statementContext) {}

// EnterConditional_expression is called when production conditional_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterConditional_expression(ctx *Conditional_expressionContext) {}

// ExitConditional_expression is called when production conditional_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitConditional_expression(ctx *Conditional_expressionContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterForever_loop_statement is called when production forever_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterForever_loop_statement(ctx *Forever_loop_statementContext) {}

// ExitForever_loop_statement is called when production forever_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitForever_loop_statement(ctx *Forever_loop_statementContext) {}

// EnterRepeat_loop_statement is called when production repeat_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterRepeat_loop_statement(ctx *Repeat_loop_statementContext) {}

// ExitRepeat_loop_statement is called when production repeat_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitRepeat_loop_statement(ctx *Repeat_loop_statementContext) {}

// EnterWhile_loop_statement is called when production while_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterWhile_loop_statement(ctx *While_loop_statementContext) {}

// ExitWhile_loop_statement is called when production while_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitWhile_loop_statement(ctx *While_loop_statementContext) {}

// EnterDo_loop_statement is called when production do_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterDo_loop_statement(ctx *Do_loop_statementContext) {}

// ExitDo_loop_statement is called when production do_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitDo_loop_statement(ctx *Do_loop_statementContext) {}

// EnterFor_loop_statement is called when production for_loop_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterFor_loop_statement(ctx *For_loop_statementContext) {}

// ExitFor_loop_statement is called when production for_loop_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitFor_loop_statement(ctx *For_loop_statementContext) {}

// EnterLoop_init_assignment is called when production loop_init_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterLoop_init_assignment(ctx *Loop_init_assignmentContext) {}

// ExitLoop_init_assignment is called when production loop_init_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitLoop_init_assignment(ctx *Loop_init_assignmentContext) {}

// EnterLoop_terminate_expression is called when production loop_terminate_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterLoop_terminate_expression(ctx *Loop_terminate_expressionContext) {
}

// ExitLoop_terminate_expression is called when production loop_terminate_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitLoop_terminate_expression(ctx *Loop_terminate_expressionContext) {
}

// EnterLoop_step_assignment is called when production loop_step_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterLoop_step_assignment(ctx *Loop_step_assignmentContext) {}

// ExitLoop_step_assignment is called when production loop_step_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitLoop_step_assignment(ctx *Loop_step_assignmentContext) {}

// EnterCase_statement is called when production case_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterCase_item_star is called when production case_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterCase_item_star(ctx *Case_item_starContext) {}

// ExitCase_item_star is called when production case_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitCase_item_star(ctx *Case_item_starContext) {}

// EnterCase_item is called when production case_item is entered.
func (s *BaseSysVerilogHDLListener) EnterCase_item(ctx *Case_itemContext) {}

// ExitCase_item is called when production case_item is exited.
func (s *BaseSysVerilogHDLListener) ExitCase_item(ctx *Case_itemContext) {}

// EnterCase_switch is called when production case_switch is entered.
func (s *BaseSysVerilogHDLListener) EnterCase_switch(ctx *Case_switchContext) {}

// ExitCase_switch is called when production case_switch is exited.
func (s *BaseSysVerilogHDLListener) ExitCase_switch(ctx *Case_switchContext) {}

// EnterCase_item_key is called when production case_item_key is entered.
func (s *BaseSysVerilogHDLListener) EnterCase_item_key(ctx *Case_item_keyContext) {}

// ExitCase_item_key is called when production case_item_key is exited.
func (s *BaseSysVerilogHDLListener) ExitCase_item_key(ctx *Case_item_keyContext) {}

// EnterCase_item_key_expression is called when production case_item_key_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterCase_item_key_expression(ctx *Case_item_key_expressionContext) {
}

// ExitCase_item_key_expression is called when production case_item_key_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitCase_item_key_expression(ctx *Case_item_key_expressionContext) {
}

// EnterComma_case_item_key_expression is called when production comma_case_item_key_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_case_item_key_expression(ctx *Comma_case_item_key_expressionContext) {
}

// ExitComma_case_item_key_expression is called when production comma_case_item_key_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_case_item_key_expression(ctx *Comma_case_item_key_expressionContext) {
}

// EnterComma_case_item_key_expression_star is called when production comma_case_item_key_expression_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_case_item_key_expression_star(ctx *Comma_case_item_key_expression_starContext) {
}

// ExitComma_case_item_key_expression_star is called when production comma_case_item_key_expression_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_case_item_key_expression_star(ctx *Comma_case_item_key_expression_starContext) {
}

// EnterExpression is called when production expression is entered.
func (s *BaseSysVerilogHDLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSysVerilogHDLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSingle_expression is called when production single_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterSingle_expression(ctx *Single_expressionContext) {}

// ExitSingle_expression is called when production single_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitSingle_expression(ctx *Single_expressionContext) {}

// EnterPrimary_range is called when production primary_range is entered.
func (s *BaseSysVerilogHDLListener) EnterPrimary_range(ctx *Primary_rangeContext) {}

// ExitPrimary_range is called when production primary_range is exited.
func (s *BaseSysVerilogHDLListener) ExitPrimary_range(ctx *Primary_rangeContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseSysVerilogHDLListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseSysVerilogHDLListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterUnary_post_assign_expression is called when production unary_post_assign_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterUnary_post_assign_expression(ctx *Unary_post_assign_expressionContext) {
}

// ExitUnary_post_assign_expression is called when production unary_post_assign_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitUnary_post_assign_expression(ctx *Unary_post_assign_expressionContext) {
}

// EnterUnary_pre_assign_expression is called when production unary_pre_assign_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterUnary_pre_assign_expression(ctx *Unary_pre_assign_expressionContext) {
}

// ExitUnary_pre_assign_expression is called when production unary_pre_assign_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitUnary_pre_assign_expression(ctx *Unary_pre_assign_expressionContext) {
}

// EnterBinary_expression is called when production binary_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterBinary_expression(ctx *Binary_expressionContext) {}

// ExitBinary_expression is called when production binary_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitBinary_expression(ctx *Binary_expressionContext) {}

// EnterTernary_expression is called when production ternary_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterTernary_expression(ctx *Ternary_expressionContext) {}

// ExitTernary_expression is called when production ternary_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitTernary_expression(ctx *Ternary_expressionContext) {}

// EnterMintypmax_expression is called when production mintypmax_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterMintypmax_expression(ctx *Mintypmax_expressionContext) {}

// ExitMintypmax_expression is called when production mintypmax_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitMintypmax_expression(ctx *Mintypmax_expressionContext) {}

// EnterStructured_value is called when production structured_value is entered.
func (s *BaseSysVerilogHDLListener) EnterStructured_value(ctx *Structured_valueContext) {}

// ExitStructured_value is called when production structured_value is exited.
func (s *BaseSysVerilogHDLListener) ExitStructured_value(ctx *Structured_valueContext) {}

// EnterArrayed_structured_value is called when production arrayed_structured_value is entered.
func (s *BaseSysVerilogHDLListener) EnterArrayed_structured_value(ctx *Arrayed_structured_valueContext) {
}

// ExitArrayed_structured_value is called when production arrayed_structured_value is exited.
func (s *BaseSysVerilogHDLListener) ExitArrayed_structured_value(ctx *Arrayed_structured_valueContext) {
}

// EnterArrayed_structure_item is called when production arrayed_structure_item is entered.
func (s *BaseSysVerilogHDLListener) EnterArrayed_structure_item(ctx *Arrayed_structure_itemContext) {}

// ExitArrayed_structure_item is called when production arrayed_structure_item is exited.
func (s *BaseSysVerilogHDLListener) ExitArrayed_structure_item(ctx *Arrayed_structure_itemContext) {}

// EnterComma_arrayed_structure_item is called when production comma_arrayed_structure_item is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_arrayed_structure_item(ctx *Comma_arrayed_structure_itemContext) {
}

// ExitComma_arrayed_structure_item is called when production comma_arrayed_structure_item is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_arrayed_structure_item(ctx *Comma_arrayed_structure_itemContext) {
}

// EnterComma_arrayed_structure_item_star is called when production comma_arrayed_structure_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_arrayed_structure_item_star(ctx *Comma_arrayed_structure_item_starContext) {
}

// ExitComma_arrayed_structure_item_star is called when production comma_arrayed_structure_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_arrayed_structure_item_star(ctx *Comma_arrayed_structure_item_starContext) {
}

// EnterArrayed_structure_item_plus is called when production arrayed_structure_item_plus is entered.
func (s *BaseSysVerilogHDLListener) EnterArrayed_structure_item_plus(ctx *Arrayed_structure_item_plusContext) {
}

// ExitArrayed_structure_item_plus is called when production arrayed_structure_item_plus is exited.
func (s *BaseSysVerilogHDLListener) ExitArrayed_structure_item_plus(ctx *Arrayed_structure_item_plusContext) {
}

// EnterVariable_type_cast is called when production variable_type_cast is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_type_cast(ctx *Variable_type_castContext) {}

// ExitVariable_type_cast is called when production variable_type_cast is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_type_cast(ctx *Variable_type_castContext) {}

// EnterWidth_type_cast is called when production width_type_cast is entered.
func (s *BaseSysVerilogHDLListener) EnterWidth_type_cast(ctx *Width_type_castContext) {}

// ExitWidth_type_cast is called when production width_type_cast is exited.
func (s *BaseSysVerilogHDLListener) ExitWidth_type_cast(ctx *Width_type_castContext) {}

// EnterSign_type_cast is called when production sign_type_cast is entered.
func (s *BaseSysVerilogHDLListener) EnterSign_type_cast(ctx *Sign_type_castContext) {}

// ExitSign_type_cast is called when production sign_type_cast is exited.
func (s *BaseSysVerilogHDLListener) ExitSign_type_cast(ctx *Sign_type_castContext) {}

// EnterNull_type_cast is called when production null_type_cast is entered.
func (s *BaseSysVerilogHDLListener) EnterNull_type_cast(ctx *Null_type_castContext) {}

// ExitNull_type_cast is called when production null_type_cast is exited.
func (s *BaseSysVerilogHDLListener) ExitNull_type_cast(ctx *Null_type_castContext) {}

// EnterVariable_type is called when production variable_type is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_type(ctx *Variable_typeContext) {}

// ExitVariable_type is called when production variable_type is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_type(ctx *Variable_typeContext) {}

// EnterType_cast_identifier is called when production type_cast_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterType_cast_identifier(ctx *Type_cast_identifierContext) {}

// ExitType_cast_identifier is called when production type_cast_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitType_cast_identifier(ctx *Type_cast_identifierContext) {}

// EnterType_cast_expression is called when production type_cast_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterType_cast_expression(ctx *Type_cast_expressionContext) {}

// ExitType_cast_expression is called when production type_cast_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitType_cast_expression(ctx *Type_cast_expressionContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterHierarchical_function_identifier is called when production hierarchical_function_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_function_identifier(ctx *Hierarchical_function_identifierContext) {
}

// ExitHierarchical_function_identifier is called when production hierarchical_function_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_function_identifier(ctx *Hierarchical_function_identifierContext) {
}

// EnterFunction_interface_assignments is called when production function_interface_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterFunction_interface_assignments(ctx *Function_interface_assignmentsContext) {
}

// ExitFunction_interface_assignments is called when production function_interface_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitFunction_interface_assignments(ctx *Function_interface_assignmentsContext) {
}

// EnterSystem_function_call is called when production system_function_call is entered.
func (s *BaseSysVerilogHDLListener) EnterSystem_function_call(ctx *System_function_callContext) {}

// ExitSystem_function_call is called when production system_function_call is exited.
func (s *BaseSysVerilogHDLListener) ExitSystem_function_call(ctx *System_function_callContext) {}

// EnterSystem_function_identifier is called when production system_function_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterSystem_function_identifier(ctx *System_function_identifierContext) {
}

// ExitSystem_function_identifier is called when production system_function_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitSystem_function_identifier(ctx *System_function_identifierContext) {
}

// EnterConstant_function_call is called when production constant_function_call is entered.
func (s *BaseSysVerilogHDLListener) EnterConstant_function_call(ctx *Constant_function_callContext) {}

// ExitConstant_function_call is called when production constant_function_call is exited.
func (s *BaseSysVerilogHDLListener) ExitConstant_function_call(ctx *Constant_function_callContext) {}

// EnterImported_function_call is called when production imported_function_call is entered.
func (s *BaseSysVerilogHDLListener) EnterImported_function_call(ctx *Imported_function_callContext) {}

// ExitImported_function_call is called when production imported_function_call is exited.
func (s *BaseSysVerilogHDLListener) ExitImported_function_call(ctx *Imported_function_callContext) {}

// EnterImported_function_hierarchical_identifier is called when production imported_function_hierarchical_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterImported_function_hierarchical_identifier(ctx *Imported_function_hierarchical_identifierContext) {
}

// ExitImported_function_hierarchical_identifier is called when production imported_function_hierarchical_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitImported_function_hierarchical_identifier(ctx *Imported_function_hierarchical_identifierContext) {
}

// EnterPrimary_hierarchical_identifier is called when production primary_hierarchical_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterPrimary_hierarchical_identifier(ctx *Primary_hierarchical_identifierContext) {
}

// ExitPrimary_hierarchical_identifier is called when production primary_hierarchical_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitPrimary_hierarchical_identifier(ctx *Primary_hierarchical_identifierContext) {
}

// EnterPrimary_imported_hierarchical_identifier is called when production primary_imported_hierarchical_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterPrimary_imported_hierarchical_identifier(ctx *Primary_imported_hierarchical_identifierContext) {
}

// ExitPrimary_imported_hierarchical_identifier is called when production primary_imported_hierarchical_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitPrimary_imported_hierarchical_identifier(ctx *Primary_imported_hierarchical_identifierContext) {
}

// EnterImported_hierarchical_identifier is called when production imported_hierarchical_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterImported_hierarchical_identifier(ctx *Imported_hierarchical_identifierContext) {
}

// ExitImported_hierarchical_identifier is called when production imported_hierarchical_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitImported_hierarchical_identifier(ctx *Imported_hierarchical_identifierContext) {
}

// EnterParenthesis_expression is called when production parenthesis_expression is entered.
func (s *BaseSysVerilogHDLListener) EnterParenthesis_expression(ctx *Parenthesis_expressionContext) {}

// ExitParenthesis_expression is called when production parenthesis_expression is exited.
func (s *BaseSysVerilogHDLListener) ExitParenthesis_expression(ctx *Parenthesis_expressionContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseSysVerilogHDLListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseSysVerilogHDLListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterMultiple_concatenation is called when production multiple_concatenation is entered.
func (s *BaseSysVerilogHDLListener) EnterMultiple_concatenation(ctx *Multiple_concatenationContext) {}

// ExitMultiple_concatenation is called when production multiple_concatenation is exited.
func (s *BaseSysVerilogHDLListener) ExitMultiple_concatenation(ctx *Multiple_concatenationContext) {}

// EnterComma_expression_plus is called when production comma_expression_plus is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_expression_plus(ctx *Comma_expression_plusContext) {}

// ExitComma_expression_plus is called when production comma_expression_plus is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_expression_plus(ctx *Comma_expression_plusContext) {}

// EnterComma_expression_star is called when production comma_expression_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_expression_star(ctx *Comma_expression_starContext) {}

// ExitComma_expression_star is called when production comma_expression_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_expression_star(ctx *Comma_expression_starContext) {}

// EnterTypedef_declaration is called when production typedef_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterTypedef_declaration(ctx *Typedef_declarationContext) {}

// ExitTypedef_declaration is called when production typedef_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitTypedef_declaration(ctx *Typedef_declarationContext) {}

// EnterTypedef_identifier is called when production typedef_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterTypedef_identifier(ctx *Typedef_identifierContext) {}

// ExitTypedef_identifier is called when production typedef_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitTypedef_identifier(ctx *Typedef_identifierContext) {}

// EnterTypedef_definition is called when production typedef_definition is entered.
func (s *BaseSysVerilogHDLListener) EnterTypedef_definition(ctx *Typedef_definitionContext) {}

// ExitTypedef_definition is called when production typedef_definition is exited.
func (s *BaseSysVerilogHDLListener) ExitTypedef_definition(ctx *Typedef_definitionContext) {}

// EnterTypedef_definition_type is called when production typedef_definition_type is entered.
func (s *BaseSysVerilogHDLListener) EnterTypedef_definition_type(ctx *Typedef_definition_typeContext) {
}

// ExitTypedef_definition_type is called when production typedef_definition_type is exited.
func (s *BaseSysVerilogHDLListener) ExitTypedef_definition_type(ctx *Typedef_definition_typeContext) {}

// EnterComplex_type is called when production complex_type is entered.
func (s *BaseSysVerilogHDLListener) EnterComplex_type(ctx *Complex_typeContext) {}

// ExitComplex_type is called when production complex_type is exited.
func (s *BaseSysVerilogHDLListener) ExitComplex_type(ctx *Complex_typeContext) {}

// EnterTypedef_type is called when production typedef_type is entered.
func (s *BaseSysVerilogHDLListener) EnterTypedef_type(ctx *Typedef_typeContext) {}

// ExitTypedef_type is called when production typedef_type is exited.
func (s *BaseSysVerilogHDLListener) ExitTypedef_type(ctx *Typedef_typeContext) {}

// EnterPar_block is called when production par_block is entered.
func (s *BaseSysVerilogHDLListener) EnterPar_block(ctx *Par_blockContext) {}

// ExitPar_block is called when production par_block is exited.
func (s *BaseSysVerilogHDLListener) ExitPar_block(ctx *Par_blockContext) {}

// EnterSeq_block is called when production seq_block is entered.
func (s *BaseSysVerilogHDLListener) EnterSeq_block(ctx *Seq_blockContext) {}

// ExitSeq_block is called when production seq_block is exited.
func (s *BaseSysVerilogHDLListener) ExitSeq_block(ctx *Seq_blockContext) {}

// EnterBlock_identifier is called when production block_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterBlock_identifier(ctx *Block_identifierContext) {}

// ExitBlock_identifier is called when production block_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitBlock_identifier(ctx *Block_identifierContext) {}

// EnterColon_block_identifier is called when production colon_block_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterColon_block_identifier(ctx *Colon_block_identifierContext) {}

// ExitColon_block_identifier is called when production colon_block_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitColon_block_identifier(ctx *Colon_block_identifierContext) {}

// EnterBlock_item_declaration_star is called when production block_item_declaration_star is entered.
func (s *BaseSysVerilogHDLListener) EnterBlock_item_declaration_star(ctx *Block_item_declaration_starContext) {
}

// ExitBlock_item_declaration_star is called when production block_item_declaration_star is exited.
func (s *BaseSysVerilogHDLListener) ExitBlock_item_declaration_star(ctx *Block_item_declaration_starContext) {
}

// EnterBlock_item_declaration_semicolon is called when production block_item_declaration_semicolon is entered.
func (s *BaseSysVerilogHDLListener) EnterBlock_item_declaration_semicolon(ctx *Block_item_declaration_semicolonContext) {
}

// ExitBlock_item_declaration_semicolon is called when production block_item_declaration_semicolon is exited.
func (s *BaseSysVerilogHDLListener) ExitBlock_item_declaration_semicolon(ctx *Block_item_declaration_semicolonContext) {
}

// EnterBlock_item_declaration is called when production block_item_declaration is entered.
func (s *BaseSysVerilogHDLListener) EnterBlock_item_declaration(ctx *Block_item_declarationContext) {}

// ExitBlock_item_declaration is called when production block_item_declaration is exited.
func (s *BaseSysVerilogHDLListener) ExitBlock_item_declaration(ctx *Block_item_declarationContext) {}

// EnterJoin_keyword is called when production join_keyword is entered.
func (s *BaseSysVerilogHDLListener) EnterJoin_keyword(ctx *Join_keywordContext) {}

// ExitJoin_keyword is called when production join_keyword is exited.
func (s *BaseSysVerilogHDLListener) ExitJoin_keyword(ctx *Join_keywordContext) {}

// EnterContinuous_assign is called when production continuous_assign is entered.
func (s *BaseSysVerilogHDLListener) EnterContinuous_assign(ctx *Continuous_assignContext) {}

// ExitContinuous_assign is called when production continuous_assign is exited.
func (s *BaseSysVerilogHDLListener) ExitContinuous_assign(ctx *Continuous_assignContext) {}

// EnterList_of_variable_assignments is called when production list_of_variable_assignments is entered.
func (s *BaseSysVerilogHDLListener) EnterList_of_variable_assignments(ctx *List_of_variable_assignmentsContext) {
}

// ExitList_of_variable_assignments is called when production list_of_variable_assignments is exited.
func (s *BaseSysVerilogHDLListener) ExitList_of_variable_assignments(ctx *List_of_variable_assignmentsContext) {
}

// EnterComma_variable_assignment_star is called when production comma_variable_assignment_star is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_variable_assignment_star(ctx *Comma_variable_assignment_starContext) {
}

// ExitComma_variable_assignment_star is called when production comma_variable_assignment_star is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_variable_assignment_star(ctx *Comma_variable_assignment_starContext) {
}

// EnterComma_variable_assignment is called when production comma_variable_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterComma_variable_assignment(ctx *Comma_variable_assignmentContext) {
}

// ExitComma_variable_assignment is called when production comma_variable_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitComma_variable_assignment(ctx *Comma_variable_assignmentContext) {
}

// EnterVariable_assignment is called when production variable_assignment is entered.
func (s *BaseSysVerilogHDLListener) EnterVariable_assignment(ctx *Variable_assignmentContext) {}

// ExitVariable_assignment is called when production variable_assignment is exited.
func (s *BaseSysVerilogHDLListener) ExitVariable_assignment(ctx *Variable_assignmentContext) {}

// EnterInitial_construct is called when production initial_construct is entered.
func (s *BaseSysVerilogHDLListener) EnterInitial_construct(ctx *Initial_constructContext) {}

// ExitInitial_construct is called when production initial_construct is exited.
func (s *BaseSysVerilogHDLListener) ExitInitial_construct(ctx *Initial_constructContext) {}

// EnterFinal_construct is called when production final_construct is entered.
func (s *BaseSysVerilogHDLListener) EnterFinal_construct(ctx *Final_constructContext) {}

// ExitFinal_construct is called when production final_construct is exited.
func (s *BaseSysVerilogHDLListener) ExitFinal_construct(ctx *Final_constructContext) {}

// EnterAlways_keyword is called when production always_keyword is entered.
func (s *BaseSysVerilogHDLListener) EnterAlways_keyword(ctx *Always_keywordContext) {}

// ExitAlways_keyword is called when production always_keyword is exited.
func (s *BaseSysVerilogHDLListener) ExitAlways_keyword(ctx *Always_keywordContext) {}

// EnterAlways_construct is called when production always_construct is entered.
func (s *BaseSysVerilogHDLListener) EnterAlways_construct(ctx *Always_constructContext) {}

// ExitAlways_construct is called when production always_construct is exited.
func (s *BaseSysVerilogHDLListener) ExitAlways_construct(ctx *Always_constructContext) {}

// EnterAttribute_instance_star is called when production attribute_instance_star is entered.
func (s *BaseSysVerilogHDLListener) EnterAttribute_instance_star(ctx *Attribute_instance_starContext) {
}

// ExitAttribute_instance_star is called when production attribute_instance_star is exited.
func (s *BaseSysVerilogHDLListener) ExitAttribute_instance_star(ctx *Attribute_instance_starContext) {}

// EnterAttribute_instance is called when production attribute_instance is entered.
func (s *BaseSysVerilogHDLListener) EnterAttribute_instance(ctx *Attribute_instanceContext) {}

// ExitAttribute_instance is called when production attribute_instance is exited.
func (s *BaseSysVerilogHDLListener) ExitAttribute_instance(ctx *Attribute_instanceContext) {}

// EnterAttr_spec_star is called when production attr_spec_star is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_spec_star(ctx *Attr_spec_starContext) {}

// ExitAttr_spec_star is called when production attr_spec_star is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_spec_star(ctx *Attr_spec_starContext) {}

// EnterAttr_spec is called when production attr_spec is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_spec(ctx *Attr_specContext) {}

// ExitAttr_spec is called when production attr_spec is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_spec(ctx *Attr_specContext) {}

// EnterAttr_name is called when production attr_name is entered.
func (s *BaseSysVerilogHDLListener) EnterAttr_name(ctx *Attr_nameContext) {}

// ExitAttr_name is called when production attr_name is exited.
func (s *BaseSysVerilogHDLListener) ExitAttr_name(ctx *Attr_nameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterHierarchical_identifier is called when production hierarchical_identifier is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_identifier(ctx *Hierarchical_identifierContext) {
}

// ExitHierarchical_identifier is called when production hierarchical_identifier is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_identifier(ctx *Hierarchical_identifierContext) {}

// EnterDot_hierarchical_identifier_branch_item_star is called when production dot_hierarchical_identifier_branch_item_star is entered.
func (s *BaseSysVerilogHDLListener) EnterDot_hierarchical_identifier_branch_item_star(ctx *Dot_hierarchical_identifier_branch_item_starContext) {
}

// ExitDot_hierarchical_identifier_branch_item_star is called when production dot_hierarchical_identifier_branch_item_star is exited.
func (s *BaseSysVerilogHDLListener) ExitDot_hierarchical_identifier_branch_item_star(ctx *Dot_hierarchical_identifier_branch_item_starContext) {
}

// EnterDot_hierarchical_identifier_branch_item is called when production dot_hierarchical_identifier_branch_item is entered.
func (s *BaseSysVerilogHDLListener) EnterDot_hierarchical_identifier_branch_item(ctx *Dot_hierarchical_identifier_branch_itemContext) {
}

// ExitDot_hierarchical_identifier_branch_item is called when production dot_hierarchical_identifier_branch_item is exited.
func (s *BaseSysVerilogHDLListener) ExitDot_hierarchical_identifier_branch_item(ctx *Dot_hierarchical_identifier_branch_itemContext) {
}

// EnterHierarchical_identifier_branch_item is called when production hierarchical_identifier_branch_item is entered.
func (s *BaseSysVerilogHDLListener) EnterHierarchical_identifier_branch_item(ctx *Hierarchical_identifier_branch_itemContext) {
}

// ExitHierarchical_identifier_branch_item is called when production hierarchical_identifier_branch_item is exited.
func (s *BaseSysVerilogHDLListener) ExitHierarchical_identifier_branch_item(ctx *Hierarchical_identifier_branch_itemContext) {
}

// EnterTimescale_compiler_directive is called when production timescale_compiler_directive is entered.
func (s *BaseSysVerilogHDLListener) EnterTimescale_compiler_directive(ctx *Timescale_compiler_directiveContext) {
}

// ExitTimescale_compiler_directive is called when production timescale_compiler_directive is exited.
func (s *BaseSysVerilogHDLListener) ExitTimescale_compiler_directive(ctx *Timescale_compiler_directiveContext) {
}

// EnterTimeunit_directive is called when production timeunit_directive is entered.
func (s *BaseSysVerilogHDLListener) EnterTimeunit_directive(ctx *Timeunit_directiveContext) {}

// ExitTimeunit_directive is called when production timeunit_directive is exited.
func (s *BaseSysVerilogHDLListener) ExitTimeunit_directive(ctx *Timeunit_directiveContext) {}

// EnterTimeprecision_directive is called when production timeprecision_directive is entered.
func (s *BaseSysVerilogHDLListener) EnterTimeprecision_directive(ctx *Timeprecision_directiveContext) {
}

// ExitTimeprecision_directive is called when production timeprecision_directive is exited.
func (s *BaseSysVerilogHDLListener) ExitTimeprecision_directive(ctx *Timeprecision_directiveContext) {}

// EnterDefault_nettype_statement is called when production default_nettype_statement is entered.
func (s *BaseSysVerilogHDLListener) EnterDefault_nettype_statement(ctx *Default_nettype_statementContext) {
}

// ExitDefault_nettype_statement is called when production default_nettype_statement is exited.
func (s *BaseSysVerilogHDLListener) ExitDefault_nettype_statement(ctx *Default_nettype_statementContext) {
}

// EnterNumber is called when production number is entered.
func (s *BaseSysVerilogHDLListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSysVerilogHDLListener) ExitNumber(ctx *NumberContext) {}

// EnterIntegral_number is called when production integral_number is entered.
func (s *BaseSysVerilogHDLListener) EnterIntegral_number(ctx *Integral_numberContext) {}

// ExitIntegral_number is called when production integral_number is exited.
func (s *BaseSysVerilogHDLListener) ExitIntegral_number(ctx *Integral_numberContext) {}

// EnterReal_number is called when production real_number is entered.
func (s *BaseSysVerilogHDLListener) EnterReal_number(ctx *Real_numberContext) {}

// ExitReal_number is called when production real_number is exited.
func (s *BaseSysVerilogHDLListener) ExitReal_number(ctx *Real_numberContext) {}
