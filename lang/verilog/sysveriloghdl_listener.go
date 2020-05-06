// Code generated from SysVerilogHDL.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SysVerilogHDL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SysVerilogHDLListener is a complete listener for a parse tree produced by SysVerilogHDLParser.
type SysVerilogHDLListener interface {
	antlr.ParseTreeListener

	// EnterModule_keyword is called when entering the module_keyword production.
	EnterModule_keyword(c *Module_keywordContext)

	// EnterStruct_keyword is called when entering the struct_keyword production.
	EnterStruct_keyword(c *Struct_keywordContext)

	// EnterAny_case_keyword is called when entering the any_case_keyword production.
	EnterAny_case_keyword(c *Any_case_keywordContext)

	// EnterSemicolon is called when entering the semicolon production.
	EnterSemicolon(c *SemicolonContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterUnary_assign_operator is called when entering the unary_assign_operator production.
	EnterUnary_assign_operator(c *Unary_assign_operatorContext)

	// EnterBinary_assign_operator is called when entering the binary_assign_operator production.
	EnterBinary_assign_operator(c *Binary_assign_operatorContext)

	// EnterSource_text is called when entering the source_text production.
	EnterSource_text(c *Source_textContext)

	// EnterDescription_star is called when entering the description_star production.
	EnterDescription_star(c *Description_starContext)

	// EnterHeader_text is called when entering the header_text production.
	EnterHeader_text(c *Header_textContext)

	// EnterDesign_attribute is called when entering the design_attribute production.
	EnterDesign_attribute(c *Design_attributeContext)

	// EnterCompiler_directive is called when entering the compiler_directive production.
	EnterCompiler_directive(c *Compiler_directiveContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterModule_declaration is called when entering the module_declaration production.
	EnterModule_declaration(c *Module_declarationContext)

	// EnterModule_identifier is called when entering the module_identifier production.
	EnterModule_identifier(c *Module_identifierContext)

	// EnterModule_interface is called when entering the module_interface production.
	EnterModule_interface(c *Module_interfaceContext)

	// EnterModule_parameter_interface is called when entering the module_parameter_interface production.
	EnterModule_parameter_interface(c *Module_parameter_interfaceContext)

	// EnterModule_port_interface is called when entering the module_port_interface production.
	EnterModule_port_interface(c *Module_port_interfaceContext)

	// EnterModule_item_star is called when entering the module_item_star production.
	EnterModule_item_star(c *Module_item_starContext)

	// EnterModule_item is called when entering the module_item production.
	EnterModule_item(c *Module_itemContext)

	// EnterColon_module_identifier is called when entering the colon_module_identifier production.
	EnterColon_module_identifier(c *Colon_module_identifierContext)

	// EnterPackage_declaration is called when entering the package_declaration production.
	EnterPackage_declaration(c *Package_declarationContext)

	// EnterPackage_identifier is called when entering the package_identifier production.
	EnterPackage_identifier(c *Package_identifierContext)

	// EnterColon_package_identifier is called when entering the colon_package_identifier production.
	EnterColon_package_identifier(c *Colon_package_identifierContext)

	// EnterPackage_item_star is called when entering the package_item_star production.
	EnterPackage_item_star(c *Package_item_starContext)

	// EnterPackage_item is called when entering the package_item production.
	EnterPackage_item(c *Package_itemContext)

	// EnterImport_package is called when entering the import_package production.
	EnterImport_package(c *Import_packageContext)

	// EnterPackage_item_identifier is called when entering the package_item_identifier production.
	EnterPackage_item_identifier(c *Package_item_identifierContext)

	// EnterParameter_item_semicolon is called when entering the parameter_item_semicolon production.
	EnterParameter_item_semicolon(c *Parameter_item_semicolonContext)

	// EnterParameter_item is called when entering the parameter_item production.
	EnterParameter_item(c *Parameter_itemContext)

	// EnterAttr_port_item_semicolon is called when entering the attr_port_item_semicolon production.
	EnterAttr_port_item_semicolon(c *Attr_port_item_semicolonContext)

	// EnterAttr_variable_item_semicolon is called when entering the attr_variable_item_semicolon production.
	EnterAttr_variable_item_semicolon(c *Attr_variable_item_semicolonContext)

	// EnterVariable_item is called when entering the variable_item production.
	EnterVariable_item(c *Variable_itemContext)

	// EnterSubroutine_item_semicolon is called when entering the subroutine_item_semicolon production.
	EnterSubroutine_item_semicolon(c *Subroutine_item_semicolonContext)

	// EnterSubroutine_item is called when entering the subroutine_item production.
	EnterSubroutine_item(c *Subroutine_itemContext)

	// EnterAttr_construct_item is called when entering the attr_construct_item production.
	EnterAttr_construct_item(c *Attr_construct_itemContext)

	// EnterConstruct_item is called when entering the construct_item production.
	EnterConstruct_item(c *Construct_itemContext)

	// EnterAttr_component_item is called when entering the attr_component_item production.
	EnterAttr_component_item(c *Attr_component_itemContext)

	// EnterComponent_item is called when entering the component_item production.
	EnterComponent_item(c *Component_itemContext)

	// EnterCompiler_item is called when entering the compiler_item production.
	EnterCompiler_item(c *Compiler_itemContext)

	// EnterType_item is called when entering the type_item production.
	EnterType_item(c *Type_itemContext)

	// EnterNull_item is called when entering the null_item production.
	EnterNull_item(c *Null_itemContext)

	// EnterList_of_interface_parameters is called when entering the list_of_interface_parameters production.
	EnterList_of_interface_parameters(c *List_of_interface_parametersContext)

	// EnterList_of_parameter_declarations is called when entering the list_of_parameter_declarations production.
	EnterList_of_parameter_declarations(c *List_of_parameter_declarationsContext)

	// EnterComma_parameter_declaration_star is called when entering the comma_parameter_declaration_star production.
	EnterComma_parameter_declaration_star(c *Comma_parameter_declaration_starContext)

	// EnterComma_parameter_declaration is called when entering the comma_parameter_declaration production.
	EnterComma_parameter_declaration(c *Comma_parameter_declarationContext)

	// EnterList_of_parameter_descriptions is called when entering the list_of_parameter_descriptions production.
	EnterList_of_parameter_descriptions(c *List_of_parameter_descriptionsContext)

	// EnterParam_declaration is called when entering the param_declaration production.
	EnterParam_declaration(c *Param_declarationContext)

	// EnterParam_description is called when entering the param_description production.
	EnterParam_description(c *Param_descriptionContext)

	// EnterParameter_declaration is called when entering the parameter_declaration production.
	EnterParameter_declaration(c *Parameter_declarationContext)

	// EnterLocal_parameter_declaration is called when entering the local_parameter_declaration production.
	EnterLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// EnterParameter_override is called when entering the parameter_override production.
	EnterParameter_override(c *Parameter_overrideContext)

	// EnterList_of_tf_interface_ports is called when entering the list_of_tf_interface_ports production.
	EnterList_of_tf_interface_ports(c *List_of_tf_interface_portsContext)

	// EnterList_of_tf_port_declarations is called when entering the list_of_tf_port_declarations production.
	EnterList_of_tf_port_declarations(c *List_of_tf_port_declarationsContext)

	// EnterList_of_tf_port_declarations_comma is called when entering the list_of_tf_port_declarations_comma production.
	EnterList_of_tf_port_declarations_comma(c *List_of_tf_port_declarations_commaContext)

	// EnterComma_attr_tf_port_declaration_star is called when entering the comma_attr_tf_port_declaration_star production.
	EnterComma_attr_tf_port_declaration_star(c *Comma_attr_tf_port_declaration_starContext)

	// EnterComma_attr_tf_port_declaration is called when entering the comma_attr_tf_port_declaration production.
	EnterComma_attr_tf_port_declaration(c *Comma_attr_tf_port_declarationContext)

	// EnterList_of_tf_port_declarations_semicolon is called when entering the list_of_tf_port_declarations_semicolon production.
	EnterList_of_tf_port_declarations_semicolon(c *List_of_tf_port_declarations_semicolonContext)

	// EnterAttr_tf_port_declaration_semicolon_plus is called when entering the attr_tf_port_declaration_semicolon_plus production.
	EnterAttr_tf_port_declaration_semicolon_plus(c *Attr_tf_port_declaration_semicolon_plusContext)

	// EnterAttr_tf_port_declaration_semicolon_star is called when entering the attr_tf_port_declaration_semicolon_star production.
	EnterAttr_tf_port_declaration_semicolon_star(c *Attr_tf_port_declaration_semicolon_starContext)

	// EnterAttr_tf_port_declaration_semicolon is called when entering the attr_tf_port_declaration_semicolon production.
	EnterAttr_tf_port_declaration_semicolon(c *Attr_tf_port_declaration_semicolonContext)

	// EnterAttr_tf_port_declaration is called when entering the attr_tf_port_declaration production.
	EnterAttr_tf_port_declaration(c *Attr_tf_port_declarationContext)

	// EnterTf_port_declaration is called when entering the tf_port_declaration production.
	EnterTf_port_declaration(c *Tf_port_declarationContext)

	// EnterList_of_interface_ports is called when entering the list_of_interface_ports production.
	EnterList_of_interface_ports(c *List_of_interface_portsContext)

	// EnterList_of_port_identifiers is called when entering the list_of_port_identifiers production.
	EnterList_of_port_identifiers(c *List_of_port_identifiersContext)

	// EnterComma_port_identifier_star is called when entering the comma_port_identifier_star production.
	EnterComma_port_identifier_star(c *Comma_port_identifier_starContext)

	// EnterComma_port_identifier is called when entering the comma_port_identifier production.
	EnterComma_port_identifier(c *Comma_port_identifierContext)

	// EnterPort_identifier is called when entering the port_identifier production.
	EnterPort_identifier(c *Port_identifierContext)

	// EnterList_of_port_declarations is called when entering the list_of_port_declarations production.
	EnterList_of_port_declarations(c *List_of_port_declarationsContext)

	// EnterList_of_port_declarations_comma is called when entering the list_of_port_declarations_comma production.
	EnterList_of_port_declarations_comma(c *List_of_port_declarations_commaContext)

	// EnterComma_attr_port_declaration_star is called when entering the comma_attr_port_declaration_star production.
	EnterComma_attr_port_declaration_star(c *Comma_attr_port_declaration_starContext)

	// EnterComma_attr_port_declaration is called when entering the comma_attr_port_declaration production.
	EnterComma_attr_port_declaration(c *Comma_attr_port_declarationContext)

	// EnterList_of_port_declarations_semicolon is called when entering the list_of_port_declarations_semicolon production.
	EnterList_of_port_declarations_semicolon(c *List_of_port_declarations_semicolonContext)

	// EnterAttr_port_declaration_semicolon_plus is called when entering the attr_port_declaration_semicolon_plus production.
	EnterAttr_port_declaration_semicolon_plus(c *Attr_port_declaration_semicolon_plusContext)

	// EnterAttr_port_declaration_semicolon_star is called when entering the attr_port_declaration_semicolon_star production.
	EnterAttr_port_declaration_semicolon_star(c *Attr_port_declaration_semicolon_starContext)

	// EnterAttr_port_declaration_semicolon is called when entering the attr_port_declaration_semicolon production.
	EnterAttr_port_declaration_semicolon(c *Attr_port_declaration_semicolonContext)

	// EnterAttr_port_declaration is called when entering the attr_port_declaration production.
	EnterAttr_port_declaration(c *Attr_port_declarationContext)

	// EnterPort_declaration is called when entering the port_declaration production.
	EnterPort_declaration(c *Port_declarationContext)

	// EnterPort_description is called when entering the port_description production.
	EnterPort_description(c *Port_descriptionContext)

	// EnterInout_description is called when entering the inout_description production.
	EnterInout_description(c *Inout_descriptionContext)

	// EnterInput_description is called when entering the input_description production.
	EnterInput_description(c *Input_descriptionContext)

	// EnterOutput_description is called when entering the output_description production.
	EnterOutput_description(c *Output_descriptionContext)

	// EnterRef_description is called when entering the ref_description production.
	EnterRef_description(c *Ref_descriptionContext)

	// EnterTf_declaration is called when entering the tf_declaration production.
	EnterTf_declaration(c *Tf_declarationContext)

	// EnterInout_declaration is called when entering the inout_declaration production.
	EnterInout_declaration(c *Inout_declarationContext)

	// EnterInput_declaration is called when entering the input_declaration production.
	EnterInput_declaration(c *Input_declarationContext)

	// EnterOutput_declaration is called when entering the output_declaration production.
	EnterOutput_declaration(c *Output_declarationContext)

	// EnterRef_declaration is called when entering the ref_declaration production.
	EnterRef_declaration(c *Ref_declarationContext)

	// EnterUser_type is called when entering the user_type production.
	EnterUser_type(c *User_typeContext)

	// EnterUser_type_identifer is called when entering the user_type_identifer production.
	EnterUser_type_identifer(c *User_type_identiferContext)

	// EnterDimension_plus is called when entering the dimension_plus production.
	EnterDimension_plus(c *Dimension_plusContext)

	// EnterDimension_star is called when entering the dimension_star production.
	EnterDimension_star(c *Dimension_starContext)

	// EnterDimension is called when entering the dimension production.
	EnterDimension(c *DimensionContext)

	// EnterRange_expression is called when entering the range_expression production.
	EnterRange_expression(c *Range_expressionContext)

	// EnterIndex_expression is called when entering the index_expression production.
	EnterIndex_expression(c *Index_expressionContext)

	// EnterSb_range is called when entering the sb_range production.
	EnterSb_range(c *Sb_rangeContext)

	// EnterBase_increment_range is called when entering the base_increment_range production.
	EnterBase_increment_range(c *Base_increment_rangeContext)

	// EnterBase_decrement_range is called when entering the base_decrement_range production.
	EnterBase_decrement_range(c *Base_decrement_rangeContext)

	// EnterBase_expression is called when entering the base_expression production.
	EnterBase_expression(c *Base_expressionContext)

	// EnterNet_type is called when entering the net_type production.
	EnterNet_type(c *Net_typeContext)

	// EnterDrive_strength is called when entering the drive_strength production.
	EnterDrive_strength(c *Drive_strengthContext)

	// EnterDrive_strength_value_0 is called when entering the drive_strength_value_0 production.
	EnterDrive_strength_value_0(c *Drive_strength_value_0Context)

	// EnterDrive_strength_value_1 is called when entering the drive_strength_value_1 production.
	EnterDrive_strength_value_1(c *Drive_strength_value_1Context)

	// EnterStrength0 is called when entering the strength0 production.
	EnterStrength0(c *Strength0Context)

	// EnterStrength1 is called when entering the strength1 production.
	EnterStrength1(c *Strength1Context)

	// EnterHighz0 is called when entering the highz0 production.
	EnterHighz0(c *Highz0Context)

	// EnterHighz1 is called when entering the highz1 production.
	EnterHighz1(c *Highz1Context)

	// EnterCharge_strength is called when entering the charge_strength production.
	EnterCharge_strength(c *Charge_strengthContext)

	// EnterCharge_size is called when entering the charge_size production.
	EnterCharge_size(c *Charge_sizeContext)

	// EnterList_of_variable_descriptions is called when entering the list_of_variable_descriptions production.
	EnterList_of_variable_descriptions(c *List_of_variable_descriptionsContext)

	// EnterComma_variable_description_star is called when entering the comma_variable_description_star production.
	EnterComma_variable_description_star(c *Comma_variable_description_starContext)

	// EnterComma_variable_description is called when entering the comma_variable_description production.
	EnterComma_variable_description(c *Comma_variable_descriptionContext)

	// EnterVariable_description is called when entering the variable_description production.
	EnterVariable_description(c *Variable_descriptionContext)

	// EnterVariable_identifier is called when entering the variable_identifier production.
	EnterVariable_identifier(c *Variable_identifierContext)

	// EnterList_of_hierarchical_variable_descriptions is called when entering the list_of_hierarchical_variable_descriptions production.
	EnterList_of_hierarchical_variable_descriptions(c *List_of_hierarchical_variable_descriptionsContext)

	// EnterComma_hierarchical_variable_description_star is called when entering the comma_hierarchical_variable_description_star production.
	EnterComma_hierarchical_variable_description_star(c *Comma_hierarchical_variable_description_starContext)

	// EnterComma_hierarchical_variable_description is called when entering the comma_hierarchical_variable_description production.
	EnterComma_hierarchical_variable_description(c *Comma_hierarchical_variable_descriptionContext)

	// EnterHierarchical_variable_description is called when entering the hierarchical_variable_description production.
	EnterHierarchical_variable_description(c *Hierarchical_variable_descriptionContext)

	// EnterHierarchical_variable_identifier is called when entering the hierarchical_variable_identifier production.
	EnterHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// EnterNet_declaration is called when entering the net_declaration production.
	EnterNet_declaration(c *Net_declarationContext)

	// EnterReg_declaration is called when entering the reg_declaration production.
	EnterReg_declaration(c *Reg_declarationContext)

	// EnterLogic_declaration is called when entering the logic_declaration production.
	EnterLogic_declaration(c *Logic_declarationContext)

	// EnterBits_type is called when entering the bits_type production.
	EnterBits_type(c *Bits_typeContext)

	// EnterBits_declaration is called when entering the bits_declaration production.
	EnterBits_declaration(c *Bits_declarationContext)

	// EnterInteger_declaration is called when entering the integer_declaration production.
	EnterInteger_declaration(c *Integer_declarationContext)

	// EnterInt_declaration is called when entering the int_declaration production.
	EnterInt_declaration(c *Int_declarationContext)

	// EnterReal_declaration is called when entering the real_declaration production.
	EnterReal_declaration(c *Real_declarationContext)

	// EnterTime_declaration is called when entering the time_declaration production.
	EnterTime_declaration(c *Time_declarationContext)

	// EnterRealtime_declaration is called when entering the realtime_declaration production.
	EnterRealtime_declaration(c *Realtime_declarationContext)

	// EnterEvent_declaration is called when entering the event_declaration production.
	EnterEvent_declaration(c *Event_declarationContext)

	// EnterGenvar_declaration is called when entering the genvar_declaration production.
	EnterGenvar_declaration(c *Genvar_declarationContext)

	// EnterUsertype_variable_declaration is called when entering the usertype_variable_declaration production.
	EnterUsertype_variable_declaration(c *Usertype_variable_declarationContext)

	// EnterString_declaration is called when entering the string_declaration production.
	EnterString_declaration(c *String_declarationContext)

	// EnterStruct_declaration is called when entering the struct_declaration production.
	EnterStruct_declaration(c *Struct_declarationContext)

	// EnterEnum_declaration is called when entering the enum_declaration production.
	EnterEnum_declaration(c *Enum_declarationContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterFunction_type is called when entering the function_type production.
	EnterFunction_type(c *Function_typeContext)

	// EnterFunction_identifier is called when entering the function_identifier production.
	EnterFunction_identifier(c *Function_identifierContext)

	// EnterFunction_interface is called when entering the function_interface production.
	EnterFunction_interface(c *Function_interfaceContext)

	// EnterFunction_item_declaration_star is called when entering the function_item_declaration_star production.
	EnterFunction_item_declaration_star(c *Function_item_declaration_starContext)

	// EnterFunction_item_declaration_semicolon is called when entering the function_item_declaration_semicolon production.
	EnterFunction_item_declaration_semicolon(c *Function_item_declaration_semicolonContext)

	// EnterFunction_item_declaration is called when entering the function_item_declaration production.
	EnterFunction_item_declaration(c *Function_item_declarationContext)

	// EnterFunction_statement is called when entering the function_statement production.
	EnterFunction_statement(c *Function_statementContext)

	// EnterColon_function_identifier is called when entering the colon_function_identifier production.
	EnterColon_function_identifier(c *Colon_function_identifierContext)

	// EnterTask_declaration is called when entering the task_declaration production.
	EnterTask_declaration(c *Task_declarationContext)

	// EnterTask_identifier is called when entering the task_identifier production.
	EnterTask_identifier(c *Task_identifierContext)

	// EnterTask_interface is called when entering the task_interface production.
	EnterTask_interface(c *Task_interfaceContext)

	// EnterTask_item_declaration_semicolon is called when entering the task_item_declaration_semicolon production.
	EnterTask_item_declaration_semicolon(c *Task_item_declaration_semicolonContext)

	// EnterTask_item_declaration is called when entering the task_item_declaration production.
	EnterTask_item_declaration(c *Task_item_declarationContext)

	// EnterTask_item_declaration_star is called when entering the task_item_declaration_star production.
	EnterTask_item_declaration_star(c *Task_item_declaration_starContext)

	// EnterTask_statement is called when entering the task_statement production.
	EnterTask_statement(c *Task_statementContext)

	// EnterStruct_item_semicolon is called when entering the struct_item_semicolon production.
	EnterStruct_item_semicolon(c *Struct_item_semicolonContext)

	// EnterStruct_item_star is called when entering the struct_item_star production.
	EnterStruct_item_star(c *Struct_item_starContext)

	// EnterStruct_item is called when entering the struct_item production.
	EnterStruct_item(c *Struct_itemContext)

	// EnterStruct_type is called when entering the struct_type production.
	EnterStruct_type(c *Struct_typeContext)

	// EnterEnum_type is called when entering the enum_type production.
	EnterEnum_type(c *Enum_typeContext)

	// EnterList_of_enum_items is called when entering the list_of_enum_items production.
	EnterList_of_enum_items(c *List_of_enum_itemsContext)

	// EnterEnum_item is called when entering the enum_item production.
	EnterEnum_item(c *Enum_itemContext)

	// EnterEnum_identifier is called when entering the enum_identifier production.
	EnterEnum_identifier(c *Enum_identifierContext)

	// EnterComma_enum_item_star is called when entering the comma_enum_item_star production.
	EnterComma_enum_item_star(c *Comma_enum_item_starContext)

	// EnterComma_enum_item is called when entering the comma_enum_item production.
	EnterComma_enum_item(c *Comma_enum_itemContext)

	// EnterEnumerated_type is called when entering the enumerated_type production.
	EnterEnumerated_type(c *Enumerated_typeContext)

	// EnterModule_instantiation is called when entering the module_instantiation production.
	EnterModule_instantiation(c *Module_instantiationContext)

	// EnterParameter_interface_assignments is called when entering the parameter_interface_assignments production.
	EnterParameter_interface_assignments(c *Parameter_interface_assignmentsContext)

	// EnterList_of_interface_assignments is called when entering the list_of_interface_assignments production.
	EnterList_of_interface_assignments(c *List_of_interface_assignmentsContext)

	// EnterList_of_ordered_interface_assignments is called when entering the list_of_ordered_interface_assignments production.
	EnterList_of_ordered_interface_assignments(c *List_of_ordered_interface_assignmentsContext)

	// EnterComma_ordered_interface_assignment_star is called when entering the comma_ordered_interface_assignment_star production.
	EnterComma_ordered_interface_assignment_star(c *Comma_ordered_interface_assignment_starContext)

	// EnterComma_ordered_interface_assignment is called when entering the comma_ordered_interface_assignment production.
	EnterComma_ordered_interface_assignment(c *Comma_ordered_interface_assignmentContext)

	// EnterOrdered_interface_assignment is called when entering the ordered_interface_assignment production.
	EnterOrdered_interface_assignment(c *Ordered_interface_assignmentContext)

	// EnterList_of_named_interface_assignments is called when entering the list_of_named_interface_assignments production.
	EnterList_of_named_interface_assignments(c *List_of_named_interface_assignmentsContext)

	// EnterComma_named_interface_assignment_star is called when entering the comma_named_interface_assignment_star production.
	EnterComma_named_interface_assignment_star(c *Comma_named_interface_assignment_starContext)

	// EnterComma_named_interface_assignment is called when entering the comma_named_interface_assignment production.
	EnterComma_named_interface_assignment(c *Comma_named_interface_assignmentContext)

	// EnterNamed_interface_assignment is called when entering the named_interface_assignment production.
	EnterNamed_interface_assignment(c *Named_interface_assignmentContext)

	// EnterList_of_module_instances is called when entering the list_of_module_instances production.
	EnterList_of_module_instances(c *List_of_module_instancesContext)

	// EnterComma_module_instance_star is called when entering the comma_module_instance_star production.
	EnterComma_module_instance_star(c *Comma_module_instance_starContext)

	// EnterComma_module_instance is called when entering the comma_module_instance production.
	EnterComma_module_instance(c *Comma_module_instanceContext)

	// EnterModule_instance is called when entering the module_instance production.
	EnterModule_instance(c *Module_instanceContext)

	// EnterModule_instance_identifier is called when entering the module_instance_identifier production.
	EnterModule_instance_identifier(c *Module_instance_identifierContext)

	// EnterArrayed_identifier is called when entering the arrayed_identifier production.
	EnterArrayed_identifier(c *Arrayed_identifierContext)

	// EnterSimple_arrayed_identifier is called when entering the simple_arrayed_identifier production.
	EnterSimple_arrayed_identifier(c *Simple_arrayed_identifierContext)

	// EnterEscaped_arrayed_identifier is called when entering the escaped_arrayed_identifier production.
	EnterEscaped_arrayed_identifier(c *Escaped_arrayed_identifierContext)

	// EnterPort_interface_assignments is called when entering the port_interface_assignments production.
	EnterPort_interface_assignments(c *Port_interface_assignmentsContext)

	// EnterDelay is called when entering the delay production.
	EnterDelay(c *DelayContext)

	// EnterList_of_delay_values is called when entering the list_of_delay_values production.
	EnterList_of_delay_values(c *List_of_delay_valuesContext)

	// EnterComma_delay_value_star is called when entering the comma_delay_value_star production.
	EnterComma_delay_value_star(c *Comma_delay_value_starContext)

	// EnterComma_delay_value is called when entering the comma_delay_value production.
	EnterComma_delay_value(c *Comma_delay_valueContext)

	// EnterDelay_value is called when entering the delay_value production.
	EnterDelay_value(c *Delay_valueContext)

	// EnterPulldown_strength is called when entering the pulldown_strength production.
	EnterPulldown_strength(c *Pulldown_strengthContext)

	// EnterPullup_strength is called when entering the pullup_strength production.
	EnterPullup_strength(c *Pullup_strengthContext)

	// EnterGate_instance_identifier is called when entering the gate_instance_identifier production.
	EnterGate_instance_identifier(c *Gate_instance_identifierContext)

	// EnterGate_instantiation is called when entering the gate_instantiation production.
	EnterGate_instantiation(c *Gate_instantiationContext)

	// EnterEnable_gatetype is called when entering the enable_gatetype production.
	EnterEnable_gatetype(c *Enable_gatetypeContext)

	// EnterMos_switchtype is called when entering the mos_switchtype production.
	EnterMos_switchtype(c *Mos_switchtypeContext)

	// EnterCmos_switchtype is called when entering the cmos_switchtype production.
	EnterCmos_switchtype(c *Cmos_switchtypeContext)

	// EnterN_output_gatetype is called when entering the n_output_gatetype production.
	EnterN_output_gatetype(c *N_output_gatetypeContext)

	// EnterN_input_gatetype is called when entering the n_input_gatetype production.
	EnterN_input_gatetype(c *N_input_gatetypeContext)

	// EnterPass_switchtype is called when entering the pass_switchtype production.
	EnterPass_switchtype(c *Pass_switchtypeContext)

	// EnterPass_enable_switchtype is called when entering the pass_enable_switchtype production.
	EnterPass_enable_switchtype(c *Pass_enable_switchtypeContext)

	// EnterPulldown_instantiation is called when entering the pulldown_instantiation production.
	EnterPulldown_instantiation(c *Pulldown_instantiationContext)

	// EnterPullup_instantiation is called when entering the pullup_instantiation production.
	EnterPullup_instantiation(c *Pullup_instantiationContext)

	// EnterEnable_instantiation is called when entering the enable_instantiation production.
	EnterEnable_instantiation(c *Enable_instantiationContext)

	// EnterMos_instantiation is called when entering the mos_instantiation production.
	EnterMos_instantiation(c *Mos_instantiationContext)

	// EnterCmos_instantiation is called when entering the cmos_instantiation production.
	EnterCmos_instantiation(c *Cmos_instantiationContext)

	// EnterN_output_instantiation is called when entering the n_output_instantiation production.
	EnterN_output_instantiation(c *N_output_instantiationContext)

	// EnterN_input_instantiation is called when entering the n_input_instantiation production.
	EnterN_input_instantiation(c *N_input_instantiationContext)

	// EnterPass_instantiation is called when entering the pass_instantiation production.
	EnterPass_instantiation(c *Pass_instantiationContext)

	// EnterPass_enable_instantiation is called when entering the pass_enable_instantiation production.
	EnterPass_enable_instantiation(c *Pass_enable_instantiationContext)

	// EnterList_of_pull_gate_instance is called when entering the list_of_pull_gate_instance production.
	EnterList_of_pull_gate_instance(c *List_of_pull_gate_instanceContext)

	// EnterList_of_enable_gate_instance is called when entering the list_of_enable_gate_instance production.
	EnterList_of_enable_gate_instance(c *List_of_enable_gate_instanceContext)

	// EnterList_of_mos_switch_instance is called when entering the list_of_mos_switch_instance production.
	EnterList_of_mos_switch_instance(c *List_of_mos_switch_instanceContext)

	// EnterList_of_cmos_switch_instance is called when entering the list_of_cmos_switch_instance production.
	EnterList_of_cmos_switch_instance(c *List_of_cmos_switch_instanceContext)

	// EnterList_of_n_input_gate_instance is called when entering the list_of_n_input_gate_instance production.
	EnterList_of_n_input_gate_instance(c *List_of_n_input_gate_instanceContext)

	// EnterList_of_n_output_gate_instance is called when entering the list_of_n_output_gate_instance production.
	EnterList_of_n_output_gate_instance(c *List_of_n_output_gate_instanceContext)

	// EnterList_of_pass_switch_instance is called when entering the list_of_pass_switch_instance production.
	EnterList_of_pass_switch_instance(c *List_of_pass_switch_instanceContext)

	// EnterList_of_pass_enable_switch_instance is called when entering the list_of_pass_enable_switch_instance production.
	EnterList_of_pass_enable_switch_instance(c *List_of_pass_enable_switch_instanceContext)

	// EnterComma_pull_gate_instance_star is called when entering the comma_pull_gate_instance_star production.
	EnterComma_pull_gate_instance_star(c *Comma_pull_gate_instance_starContext)

	// EnterComma_enable_gate_instance_star is called when entering the comma_enable_gate_instance_star production.
	EnterComma_enable_gate_instance_star(c *Comma_enable_gate_instance_starContext)

	// EnterComma_mos_switch_instance_star is called when entering the comma_mos_switch_instance_star production.
	EnterComma_mos_switch_instance_star(c *Comma_mos_switch_instance_starContext)

	// EnterComma_cmos_switch_instance_star is called when entering the comma_cmos_switch_instance_star production.
	EnterComma_cmos_switch_instance_star(c *Comma_cmos_switch_instance_starContext)

	// EnterComma_n_input_gate_instance_star is called when entering the comma_n_input_gate_instance_star production.
	EnterComma_n_input_gate_instance_star(c *Comma_n_input_gate_instance_starContext)

	// EnterComma_n_output_gate_instance_star is called when entering the comma_n_output_gate_instance_star production.
	EnterComma_n_output_gate_instance_star(c *Comma_n_output_gate_instance_starContext)

	// EnterComma_pass_switch_instance_star is called when entering the comma_pass_switch_instance_star production.
	EnterComma_pass_switch_instance_star(c *Comma_pass_switch_instance_starContext)

	// EnterComma_pass_enable_switch_instance_star is called when entering the comma_pass_enable_switch_instance_star production.
	EnterComma_pass_enable_switch_instance_star(c *Comma_pass_enable_switch_instance_starContext)

	// EnterComma_pull_gate_instance is called when entering the comma_pull_gate_instance production.
	EnterComma_pull_gate_instance(c *Comma_pull_gate_instanceContext)

	// EnterComma_enable_gate_instance is called when entering the comma_enable_gate_instance production.
	EnterComma_enable_gate_instance(c *Comma_enable_gate_instanceContext)

	// EnterComma_mos_switch_instance is called when entering the comma_mos_switch_instance production.
	EnterComma_mos_switch_instance(c *Comma_mos_switch_instanceContext)

	// EnterComma_cmos_switch_instance is called when entering the comma_cmos_switch_instance production.
	EnterComma_cmos_switch_instance(c *Comma_cmos_switch_instanceContext)

	// EnterComma_n_input_gate_instance is called when entering the comma_n_input_gate_instance production.
	EnterComma_n_input_gate_instance(c *Comma_n_input_gate_instanceContext)

	// EnterComma_n_output_gate_instance is called when entering the comma_n_output_gate_instance production.
	EnterComma_n_output_gate_instance(c *Comma_n_output_gate_instanceContext)

	// EnterComma_pass_switch_instance is called when entering the comma_pass_switch_instance production.
	EnterComma_pass_switch_instance(c *Comma_pass_switch_instanceContext)

	// EnterComma_pass_enable_switch_instance is called when entering the comma_pass_enable_switch_instance production.
	EnterComma_pass_enable_switch_instance(c *Comma_pass_enable_switch_instanceContext)

	// EnterPull_gate_instance is called when entering the pull_gate_instance production.
	EnterPull_gate_instance(c *Pull_gate_instanceContext)

	// EnterEnable_gate_instance is called when entering the enable_gate_instance production.
	EnterEnable_gate_instance(c *Enable_gate_instanceContext)

	// EnterMos_switch_instance is called when entering the mos_switch_instance production.
	EnterMos_switch_instance(c *Mos_switch_instanceContext)

	// EnterCmos_switch_instance is called when entering the cmos_switch_instance production.
	EnterCmos_switch_instance(c *Cmos_switch_instanceContext)

	// EnterN_input_gate_instance is called when entering the n_input_gate_instance production.
	EnterN_input_gate_instance(c *N_input_gate_instanceContext)

	// EnterN_output_gate_instance is called when entering the n_output_gate_instance production.
	EnterN_output_gate_instance(c *N_output_gate_instanceContext)

	// EnterPass_switch_instance is called when entering the pass_switch_instance production.
	EnterPass_switch_instance(c *Pass_switch_instanceContext)

	// EnterPass_enable_switch_instance is called when entering the pass_enable_switch_instance production.
	EnterPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// EnterPull_gate_interface is called when entering the pull_gate_interface production.
	EnterPull_gate_interface(c *Pull_gate_interfaceContext)

	// EnterEnable_gate_interface is called when entering the enable_gate_interface production.
	EnterEnable_gate_interface(c *Enable_gate_interfaceContext)

	// EnterMos_switch_interface is called when entering the mos_switch_interface production.
	EnterMos_switch_interface(c *Mos_switch_interfaceContext)

	// EnterCmos_switch_interface is called when entering the cmos_switch_interface production.
	EnterCmos_switch_interface(c *Cmos_switch_interfaceContext)

	// EnterN_input_gate_interface is called when entering the n_input_gate_interface production.
	EnterN_input_gate_interface(c *N_input_gate_interfaceContext)

	// EnterN_output_gate_interface is called when entering the n_output_gate_interface production.
	EnterN_output_gate_interface(c *N_output_gate_interfaceContext)

	// EnterPass_switch_interface is called when entering the pass_switch_interface production.
	EnterPass_switch_interface(c *Pass_switch_interfaceContext)

	// EnterPass_enable_switch_interface is called when entering the pass_enable_switch_interface production.
	EnterPass_enable_switch_interface(c *Pass_enable_switch_interfaceContext)

	// EnterList_of_input_terminals is called when entering the list_of_input_terminals production.
	EnterList_of_input_terminals(c *List_of_input_terminalsContext)

	// EnterList_of_output_terminals is called when entering the list_of_output_terminals production.
	EnterList_of_output_terminals(c *List_of_output_terminalsContext)

	// EnterComma_input_terminal_star is called when entering the comma_input_terminal_star production.
	EnterComma_input_terminal_star(c *Comma_input_terminal_starContext)

	// EnterComma_output_terminal_star is called when entering the comma_output_terminal_star production.
	EnterComma_output_terminal_star(c *Comma_output_terminal_starContext)

	// EnterComma_input_terminal is called when entering the comma_input_terminal production.
	EnterComma_input_terminal(c *Comma_input_terminalContext)

	// EnterComma_output_terminal is called when entering the comma_output_terminal production.
	EnterComma_output_terminal(c *Comma_output_terminalContext)

	// EnterEnable_terminal is called when entering the enable_terminal production.
	EnterEnable_terminal(c *Enable_terminalContext)

	// EnterInput_terminal is called when entering the input_terminal production.
	EnterInput_terminal(c *Input_terminalContext)

	// EnterInout_terminal is called when entering the inout_terminal production.
	EnterInout_terminal(c *Inout_terminalContext)

	// EnterNcontrol_terminal is called when entering the ncontrol_terminal production.
	EnterNcontrol_terminal(c *Ncontrol_terminalContext)

	// EnterOutput_terminal is called when entering the output_terminal production.
	EnterOutput_terminal(c *Output_terminalContext)

	// EnterPcontrol_terminal is called when entering the pcontrol_terminal production.
	EnterPcontrol_terminal(c *Pcontrol_terminalContext)

	// EnterStatement_star is called when entering the statement_star production.
	EnterStatement_star(c *Statement_starContext)

	// EnterStatement_semicolon is called when entering the statement_semicolon production.
	EnterStatement_semicolon(c *Statement_semicolonContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment_statement is called when entering the assignment_statement production.
	EnterAssignment_statement(c *Assignment_statementContext)

	// EnterFlow_control_statement is called when entering the flow_control_statement production.
	EnterFlow_control_statement(c *Flow_control_statementContext)

	// EnterBlock_statement is called when entering the block_statement production.
	EnterBlock_statement(c *Block_statementContext)

	// EnterTask_call_statement is called when entering the task_call_statement production.
	EnterTask_call_statement(c *Task_call_statementContext)

	// EnterEvent_statement is called when entering the event_statement production.
	EnterEvent_statement(c *Event_statementContext)

	// EnterProcedural_statement is called when entering the procedural_statement production.
	EnterProcedural_statement(c *Procedural_statementContext)

	// EnterExpression_statement is called when entering the expression_statement production.
	EnterExpression_statement(c *Expression_statementContext)

	// EnterSubroutine_statement is called when entering the subroutine_statement production.
	EnterSubroutine_statement(c *Subroutine_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterNull_statement is called when entering the null_statement production.
	EnterNull_statement(c *Null_statementContext)

	// EnterProcedural_continuous_assignments is called when entering the procedural_continuous_assignments production.
	EnterProcedural_continuous_assignments(c *Procedural_continuous_assignmentsContext)

	// EnterAssign_statement is called when entering the assign_statement production.
	EnterAssign_statement(c *Assign_statementContext)

	// EnterDeassign_statement is called when entering the deassign_statement production.
	EnterDeassign_statement(c *Deassign_statementContext)

	// EnterForce_statement is called when entering the force_statement production.
	EnterForce_statement(c *Force_statementContext)

	// EnterRelease_statement is called when entering the release_statement production.
	EnterRelease_statement(c *Release_statementContext)

	// EnterProcedural_timing_control_statement is called when entering the procedural_timing_control_statement production.
	EnterProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// EnterProperty_statement is called when entering the property_statement production.
	EnterProperty_statement(c *Property_statementContext)

	// EnterDisable_condition_statement is called when entering the disable_condition_statement production.
	EnterDisable_condition_statement(c *Disable_condition_statementContext)

	// EnterProperty_expression is called when entering the property_expression production.
	EnterProperty_expression(c *Property_expressionContext)

	// EnterProcedural_assertion_statement is called when entering the procedural_assertion_statement production.
	EnterProcedural_assertion_statement(c *Procedural_assertion_statementContext)

	// EnterAssert_else_statement is called when entering the assert_else_statement production.
	EnterAssert_else_statement(c *Assert_else_statementContext)

	// EnterAssert_statement is called when entering the assert_statement production.
	EnterAssert_statement(c *Assert_statementContext)

	// EnterSystem_task_enable is called when entering the system_task_enable production.
	EnterSystem_task_enable(c *System_task_enableContext)

	// EnterSystem_task_identifier is called when entering the system_task_identifier production.
	EnterSystem_task_identifier(c *System_task_identifierContext)

	// EnterTask_interface_assignments is called when entering the task_interface_assignments production.
	EnterTask_interface_assignments(c *Task_interface_assignmentsContext)

	// EnterTask_enable is called when entering the task_enable production.
	EnterTask_enable(c *Task_enableContext)

	// EnterHierarchical_task_identifier is called when entering the hierarchical_task_identifier production.
	EnterHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// EnterDisable_statement is called when entering the disable_statement production.
	EnterDisable_statement(c *Disable_statementContext)

	// EnterHierarchical_block_identifier is called when entering the hierarchical_block_identifier production.
	EnterHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// EnterVariable_lvalue is called when entering the variable_lvalue production.
	EnterVariable_lvalue(c *Variable_lvalueContext)

	// EnterHierarchical_variable_lvalue is called when entering the hierarchical_variable_lvalue production.
	EnterHierarchical_variable_lvalue(c *Hierarchical_variable_lvalueContext)

	// EnterVariable_concatenation is called when entering the variable_concatenation production.
	EnterVariable_concatenation(c *Variable_concatenationContext)

	// EnterVariable_concatenation_value is called when entering the variable_concatenation_value production.
	EnterVariable_concatenation_value(c *Variable_concatenation_valueContext)

	// EnterComma_vcv_star is called when entering the comma_vcv_star production.
	EnterComma_vcv_star(c *Comma_vcv_starContext)

	// EnterBlocking_assignment is called when entering the blocking_assignment production.
	EnterBlocking_assignment(c *Blocking_assignmentContext)

	// EnterNonblocking_assignment is called when entering the nonblocking_assignment production.
	EnterNonblocking_assignment(c *Nonblocking_assignmentContext)

	// EnterPrefix_assignment is called when entering the prefix_assignment production.
	EnterPrefix_assignment(c *Prefix_assignmentContext)

	// EnterPostfix_assignment is called when entering the postfix_assignment production.
	EnterPostfix_assignment(c *Postfix_assignmentContext)

	// EnterOperator_assignment is called when entering the operator_assignment production.
	EnterOperator_assignment(c *Operator_assignmentContext)

	// EnterDeclarative_assignment is called when entering the declarative_assignment production.
	EnterDeclarative_assignment(c *Declarative_assignmentContext)

	// EnterDelay_or_event_control is called when entering the delay_or_event_control production.
	EnterDelay_or_event_control(c *Delay_or_event_controlContext)

	// EnterDelay_control is called when entering the delay_control production.
	EnterDelay_control(c *Delay_controlContext)

	// EnterEvent_control is called when entering the event_control production.
	EnterEvent_control(c *Event_controlContext)

	// EnterEvent_control_identifier is called when entering the event_control_identifier production.
	EnterEvent_control_identifier(c *Event_control_identifierContext)

	// EnterEvent_control_expression is called when entering the event_control_expression production.
	EnterEvent_control_expression(c *Event_control_expressionContext)

	// EnterEvent_expression is called when entering the event_expression production.
	EnterEvent_expression(c *Event_expressionContext)

	// EnterSingle_event_expression is called when entering the single_event_expression production.
	EnterSingle_event_expression(c *Single_event_expressionContext)

	// EnterEvent_expression_edgespec is called when entering the event_expression_edgespec production.
	EnterEvent_expression_edgespec(c *Event_expression_edgespecContext)

	// EnterEvent_expression_or is called when entering the event_expression_or production.
	EnterEvent_expression_or(c *Event_expression_orContext)

	// EnterList_of_event_expression_comma is called when entering the list_of_event_expression_comma production.
	EnterList_of_event_expression_comma(c *List_of_event_expression_commaContext)

	// EnterComma_event_expression_star is called when entering the comma_event_expression_star production.
	EnterComma_event_expression_star(c *Comma_event_expression_starContext)

	// EnterComma_event_expression is called when entering the comma_event_expression production.
	EnterComma_event_expression(c *Comma_event_expressionContext)

	// EnterList_of_event_expression_or is called when entering the list_of_event_expression_or production.
	EnterList_of_event_expression_or(c *List_of_event_expression_orContext)

	// EnterOr_event_expression_star is called when entering the or_event_expression_star production.
	EnterOr_event_expression_star(c *Or_event_expression_starContext)

	// EnterOr_event_expression is called when entering the or_event_expression production.
	EnterOr_event_expression(c *Or_event_expressionContext)

	// EnterEvent_control_wildcard is called when entering the event_control_wildcard production.
	EnterEvent_control_wildcard(c *Event_control_wildcardContext)

	// EnterRepeat_event_control is called when entering the repeat_event_control production.
	EnterRepeat_event_control(c *Repeat_event_controlContext)

	// EnterEvent_trigger is called when entering the event_trigger production.
	EnterEvent_trigger(c *Event_triggerContext)

	// EnterHierarchical_event_identifier is called when entering the hierarchical_event_identifier production.
	EnterHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// EnterEvent_identifier is called when entering the event_identifier production.
	EnterEvent_identifier(c *Event_identifierContext)

	// EnterWait_statement is called when entering the wait_statement production.
	EnterWait_statement(c *Wait_statementContext)

	// EnterAttr_generated_instantiation is called when entering the attr_generated_instantiation production.
	EnterAttr_generated_instantiation(c *Attr_generated_instantiationContext)

	// EnterGenerated_instantiation is called when entering the generated_instantiation production.
	EnterGenerated_instantiation(c *Generated_instantiationContext)

	// EnterGenerate_item_star is called when entering the generate_item_star production.
	EnterGenerate_item_star(c *Generate_item_starContext)

	// EnterGenerate_item is called when entering the generate_item production.
	EnterGenerate_item(c *Generate_itemContext)

	// EnterGenerate_block is called when entering the generate_block production.
	EnterGenerate_block(c *Generate_blockContext)

	// EnterGenerate_colon_block_identifier0 is called when entering the generate_colon_block_identifier0 production.
	EnterGenerate_colon_block_identifier0(c *Generate_colon_block_identifier0Context)

	// EnterGenerate_colon_block_identifier1 is called when entering the generate_colon_block_identifier1 production.
	EnterGenerate_colon_block_identifier1(c *Generate_colon_block_identifier1Context)

	// EnterGenerate_colon_block_identifier is called when entering the generate_colon_block_identifier production.
	EnterGenerate_colon_block_identifier(c *Generate_colon_block_identifierContext)

	// EnterGenerate_block_identifier is called when entering the generate_block_identifier production.
	EnterGenerate_block_identifier(c *Generate_block_identifierContext)

	// EnterGenerate_conditional_statement is called when entering the generate_conditional_statement production.
	EnterGenerate_conditional_statement(c *Generate_conditional_statementContext)

	// EnterGenerate_if_statement is called when entering the generate_if_statement production.
	EnterGenerate_if_statement(c *Generate_if_statementContext)

	// EnterGenerate_else_statement is called when entering the generate_else_statement production.
	EnterGenerate_else_statement(c *Generate_else_statementContext)

	// EnterGenerate_loop_statement is called when entering the generate_loop_statement production.
	EnterGenerate_loop_statement(c *Generate_loop_statementContext)

	// EnterGenerate_forever_loop_statement is called when entering the generate_forever_loop_statement production.
	EnterGenerate_forever_loop_statement(c *Generate_forever_loop_statementContext)

	// EnterGenerate_repeat_loop_statement is called when entering the generate_repeat_loop_statement production.
	EnterGenerate_repeat_loop_statement(c *Generate_repeat_loop_statementContext)

	// EnterGenerate_while_loop_statement is called when entering the generate_while_loop_statement production.
	EnterGenerate_while_loop_statement(c *Generate_while_loop_statementContext)

	// EnterGenerate_do_loop_statement is called when entering the generate_do_loop_statement production.
	EnterGenerate_do_loop_statement(c *Generate_do_loop_statementContext)

	// EnterGenerate_for_loop_statement is called when entering the generate_for_loop_statement production.
	EnterGenerate_for_loop_statement(c *Generate_for_loop_statementContext)

	// EnterGenerate_case_statement is called when entering the generate_case_statement production.
	EnterGenerate_case_statement(c *Generate_case_statementContext)

	// EnterGenerate_case_item_star is called when entering the generate_case_item_star production.
	EnterGenerate_case_item_star(c *Generate_case_item_starContext)

	// EnterGenerate_case_item is called when entering the generate_case_item production.
	EnterGenerate_case_item(c *Generate_case_itemContext)

	// EnterConditional_statement is called when entering the conditional_statement production.
	EnterConditional_statement(c *Conditional_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterElse_statement is called when entering the else_statement production.
	EnterElse_statement(c *Else_statementContext)

	// EnterConditional_expression is called when entering the conditional_expression production.
	EnterConditional_expression(c *Conditional_expressionContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterForever_loop_statement is called when entering the forever_loop_statement production.
	EnterForever_loop_statement(c *Forever_loop_statementContext)

	// EnterRepeat_loop_statement is called when entering the repeat_loop_statement production.
	EnterRepeat_loop_statement(c *Repeat_loop_statementContext)

	// EnterWhile_loop_statement is called when entering the while_loop_statement production.
	EnterWhile_loop_statement(c *While_loop_statementContext)

	// EnterDo_loop_statement is called when entering the do_loop_statement production.
	EnterDo_loop_statement(c *Do_loop_statementContext)

	// EnterFor_loop_statement is called when entering the for_loop_statement production.
	EnterFor_loop_statement(c *For_loop_statementContext)

	// EnterLoop_init_assignment is called when entering the loop_init_assignment production.
	EnterLoop_init_assignment(c *Loop_init_assignmentContext)

	// EnterLoop_terminate_expression is called when entering the loop_terminate_expression production.
	EnterLoop_terminate_expression(c *Loop_terminate_expressionContext)

	// EnterLoop_step_assignment is called when entering the loop_step_assignment production.
	EnterLoop_step_assignment(c *Loop_step_assignmentContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterCase_item_star is called when entering the case_item_star production.
	EnterCase_item_star(c *Case_item_starContext)

	// EnterCase_item is called when entering the case_item production.
	EnterCase_item(c *Case_itemContext)

	// EnterCase_switch is called when entering the case_switch production.
	EnterCase_switch(c *Case_switchContext)

	// EnterCase_item_key is called when entering the case_item_key production.
	EnterCase_item_key(c *Case_item_keyContext)

	// EnterCase_item_key_expression is called when entering the case_item_key_expression production.
	EnterCase_item_key_expression(c *Case_item_key_expressionContext)

	// EnterComma_case_item_key_expression is called when entering the comma_case_item_key_expression production.
	EnterComma_case_item_key_expression(c *Comma_case_item_key_expressionContext)

	// EnterComma_case_item_key_expression_star is called when entering the comma_case_item_key_expression_star production.
	EnterComma_case_item_key_expression_star(c *Comma_case_item_key_expression_starContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSingle_expression is called when entering the single_expression production.
	EnterSingle_expression(c *Single_expressionContext)

	// EnterPrimary_range is called when entering the primary_range production.
	EnterPrimary_range(c *Primary_rangeContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterUnary_expression is called when entering the unary_expression production.
	EnterUnary_expression(c *Unary_expressionContext)

	// EnterUnary_post_assign_expression is called when entering the unary_post_assign_expression production.
	EnterUnary_post_assign_expression(c *Unary_post_assign_expressionContext)

	// EnterUnary_pre_assign_expression is called when entering the unary_pre_assign_expression production.
	EnterUnary_pre_assign_expression(c *Unary_pre_assign_expressionContext)

	// EnterBinary_expression is called when entering the binary_expression production.
	EnterBinary_expression(c *Binary_expressionContext)

	// EnterTernary_expression is called when entering the ternary_expression production.
	EnterTernary_expression(c *Ternary_expressionContext)

	// EnterMintypmax_expression is called when entering the mintypmax_expression production.
	EnterMintypmax_expression(c *Mintypmax_expressionContext)

	// EnterStructured_value is called when entering the structured_value production.
	EnterStructured_value(c *Structured_valueContext)

	// EnterArrayed_structured_value is called when entering the arrayed_structured_value production.
	EnterArrayed_structured_value(c *Arrayed_structured_valueContext)

	// EnterArrayed_structure_item is called when entering the arrayed_structure_item production.
	EnterArrayed_structure_item(c *Arrayed_structure_itemContext)

	// EnterComma_arrayed_structure_item is called when entering the comma_arrayed_structure_item production.
	EnterComma_arrayed_structure_item(c *Comma_arrayed_structure_itemContext)

	// EnterComma_arrayed_structure_item_star is called when entering the comma_arrayed_structure_item_star production.
	EnterComma_arrayed_structure_item_star(c *Comma_arrayed_structure_item_starContext)

	// EnterArrayed_structure_item_plus is called when entering the arrayed_structure_item_plus production.
	EnterArrayed_structure_item_plus(c *Arrayed_structure_item_plusContext)

	// EnterVariable_type_cast is called when entering the variable_type_cast production.
	EnterVariable_type_cast(c *Variable_type_castContext)

	// EnterWidth_type_cast is called when entering the width_type_cast production.
	EnterWidth_type_cast(c *Width_type_castContext)

	// EnterSign_type_cast is called when entering the sign_type_cast production.
	EnterSign_type_cast(c *Sign_type_castContext)

	// EnterNull_type_cast is called when entering the null_type_cast production.
	EnterNull_type_cast(c *Null_type_castContext)

	// EnterVariable_type is called when entering the variable_type production.
	EnterVariable_type(c *Variable_typeContext)

	// EnterType_cast_identifier is called when entering the type_cast_identifier production.
	EnterType_cast_identifier(c *Type_cast_identifierContext)

	// EnterType_cast_expression is called when entering the type_cast_expression production.
	EnterType_cast_expression(c *Type_cast_expressionContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterHierarchical_function_identifier is called when entering the hierarchical_function_identifier production.
	EnterHierarchical_function_identifier(c *Hierarchical_function_identifierContext)

	// EnterFunction_interface_assignments is called when entering the function_interface_assignments production.
	EnterFunction_interface_assignments(c *Function_interface_assignmentsContext)

	// EnterSystem_function_call is called when entering the system_function_call production.
	EnterSystem_function_call(c *System_function_callContext)

	// EnterSystem_function_identifier is called when entering the system_function_identifier production.
	EnterSystem_function_identifier(c *System_function_identifierContext)

	// EnterConstant_function_call is called when entering the constant_function_call production.
	EnterConstant_function_call(c *Constant_function_callContext)

	// EnterImported_function_call is called when entering the imported_function_call production.
	EnterImported_function_call(c *Imported_function_callContext)

	// EnterImported_function_hierarchical_identifier is called when entering the imported_function_hierarchical_identifier production.
	EnterImported_function_hierarchical_identifier(c *Imported_function_hierarchical_identifierContext)

	// EnterPrimary_hierarchical_identifier is called when entering the primary_hierarchical_identifier production.
	EnterPrimary_hierarchical_identifier(c *Primary_hierarchical_identifierContext)

	// EnterPrimary_imported_hierarchical_identifier is called when entering the primary_imported_hierarchical_identifier production.
	EnterPrimary_imported_hierarchical_identifier(c *Primary_imported_hierarchical_identifierContext)

	// EnterImported_hierarchical_identifier is called when entering the imported_hierarchical_identifier production.
	EnterImported_hierarchical_identifier(c *Imported_hierarchical_identifierContext)

	// EnterParenthesis_expression is called when entering the parenthesis_expression production.
	EnterParenthesis_expression(c *Parenthesis_expressionContext)

	// EnterConcatenation is called when entering the concatenation production.
	EnterConcatenation(c *ConcatenationContext)

	// EnterMultiple_concatenation is called when entering the multiple_concatenation production.
	EnterMultiple_concatenation(c *Multiple_concatenationContext)

	// EnterComma_expression_plus is called when entering the comma_expression_plus production.
	EnterComma_expression_plus(c *Comma_expression_plusContext)

	// EnterComma_expression_star is called when entering the comma_expression_star production.
	EnterComma_expression_star(c *Comma_expression_starContext)

	// EnterTypedef_declaration is called when entering the typedef_declaration production.
	EnterTypedef_declaration(c *Typedef_declarationContext)

	// EnterTypedef_identifier is called when entering the typedef_identifier production.
	EnterTypedef_identifier(c *Typedef_identifierContext)

	// EnterTypedef_definition is called when entering the typedef_definition production.
	EnterTypedef_definition(c *Typedef_definitionContext)

	// EnterTypedef_definition_type is called when entering the typedef_definition_type production.
	EnterTypedef_definition_type(c *Typedef_definition_typeContext)

	// EnterComplex_type is called when entering the complex_type production.
	EnterComplex_type(c *Complex_typeContext)

	// EnterTypedef_type is called when entering the typedef_type production.
	EnterTypedef_type(c *Typedef_typeContext)

	// EnterPar_block is called when entering the par_block production.
	EnterPar_block(c *Par_blockContext)

	// EnterSeq_block is called when entering the seq_block production.
	EnterSeq_block(c *Seq_blockContext)

	// EnterBlock_identifier is called when entering the block_identifier production.
	EnterBlock_identifier(c *Block_identifierContext)

	// EnterColon_block_identifier is called when entering the colon_block_identifier production.
	EnterColon_block_identifier(c *Colon_block_identifierContext)

	// EnterBlock_item_declaration_star is called when entering the block_item_declaration_star production.
	EnterBlock_item_declaration_star(c *Block_item_declaration_starContext)

	// EnterBlock_item_declaration_semicolon is called when entering the block_item_declaration_semicolon production.
	EnterBlock_item_declaration_semicolon(c *Block_item_declaration_semicolonContext)

	// EnterBlock_item_declaration is called when entering the block_item_declaration production.
	EnterBlock_item_declaration(c *Block_item_declarationContext)

	// EnterJoin_keyword is called when entering the join_keyword production.
	EnterJoin_keyword(c *Join_keywordContext)

	// EnterContinuous_assign is called when entering the continuous_assign production.
	EnterContinuous_assign(c *Continuous_assignContext)

	// EnterList_of_variable_assignments is called when entering the list_of_variable_assignments production.
	EnterList_of_variable_assignments(c *List_of_variable_assignmentsContext)

	// EnterComma_variable_assignment_star is called when entering the comma_variable_assignment_star production.
	EnterComma_variable_assignment_star(c *Comma_variable_assignment_starContext)

	// EnterComma_variable_assignment is called when entering the comma_variable_assignment production.
	EnterComma_variable_assignment(c *Comma_variable_assignmentContext)

	// EnterVariable_assignment is called when entering the variable_assignment production.
	EnterVariable_assignment(c *Variable_assignmentContext)

	// EnterInitial_construct is called when entering the initial_construct production.
	EnterInitial_construct(c *Initial_constructContext)

	// EnterFinal_construct is called when entering the final_construct production.
	EnterFinal_construct(c *Final_constructContext)

	// EnterAlways_keyword is called when entering the always_keyword production.
	EnterAlways_keyword(c *Always_keywordContext)

	// EnterAlways_construct is called when entering the always_construct production.
	EnterAlways_construct(c *Always_constructContext)

	// EnterAttribute_instance_star is called when entering the attribute_instance_star production.
	EnterAttribute_instance_star(c *Attribute_instance_starContext)

	// EnterAttribute_instance is called when entering the attribute_instance production.
	EnterAttribute_instance(c *Attribute_instanceContext)

	// EnterAttr_spec_star is called when entering the attr_spec_star production.
	EnterAttr_spec_star(c *Attr_spec_starContext)

	// EnterAttr_spec is called when entering the attr_spec production.
	EnterAttr_spec(c *Attr_specContext)

	// EnterAttr_name is called when entering the attr_name production.
	EnterAttr_name(c *Attr_nameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterHierarchical_identifier is called when entering the hierarchical_identifier production.
	EnterHierarchical_identifier(c *Hierarchical_identifierContext)

	// EnterDot_hierarchical_identifier_branch_item_star is called when entering the dot_hierarchical_identifier_branch_item_star production.
	EnterDot_hierarchical_identifier_branch_item_star(c *Dot_hierarchical_identifier_branch_item_starContext)

	// EnterDot_hierarchical_identifier_branch_item is called when entering the dot_hierarchical_identifier_branch_item production.
	EnterDot_hierarchical_identifier_branch_item(c *Dot_hierarchical_identifier_branch_itemContext)

	// EnterHierarchical_identifier_branch_item is called when entering the hierarchical_identifier_branch_item production.
	EnterHierarchical_identifier_branch_item(c *Hierarchical_identifier_branch_itemContext)

	// EnterTimescale_compiler_directive is called when entering the timescale_compiler_directive production.
	EnterTimescale_compiler_directive(c *Timescale_compiler_directiveContext)

	// EnterTimeunit_directive is called when entering the timeunit_directive production.
	EnterTimeunit_directive(c *Timeunit_directiveContext)

	// EnterTimeprecision_directive is called when entering the timeprecision_directive production.
	EnterTimeprecision_directive(c *Timeprecision_directiveContext)

	// EnterDefault_nettype_statement is called when entering the default_nettype_statement production.
	EnterDefault_nettype_statement(c *Default_nettype_statementContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterIntegral_number is called when entering the integral_number production.
	EnterIntegral_number(c *Integral_numberContext)

	// EnterReal_number is called when entering the real_number production.
	EnterReal_number(c *Real_numberContext)

	// ExitModule_keyword is called when exiting the module_keyword production.
	ExitModule_keyword(c *Module_keywordContext)

	// ExitStruct_keyword is called when exiting the struct_keyword production.
	ExitStruct_keyword(c *Struct_keywordContext)

	// ExitAny_case_keyword is called when exiting the any_case_keyword production.
	ExitAny_case_keyword(c *Any_case_keywordContext)

	// ExitSemicolon is called when exiting the semicolon production.
	ExitSemicolon(c *SemicolonContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitUnary_assign_operator is called when exiting the unary_assign_operator production.
	ExitUnary_assign_operator(c *Unary_assign_operatorContext)

	// ExitBinary_assign_operator is called when exiting the binary_assign_operator production.
	ExitBinary_assign_operator(c *Binary_assign_operatorContext)

	// ExitSource_text is called when exiting the source_text production.
	ExitSource_text(c *Source_textContext)

	// ExitDescription_star is called when exiting the description_star production.
	ExitDescription_star(c *Description_starContext)

	// ExitHeader_text is called when exiting the header_text production.
	ExitHeader_text(c *Header_textContext)

	// ExitDesign_attribute is called when exiting the design_attribute production.
	ExitDesign_attribute(c *Design_attributeContext)

	// ExitCompiler_directive is called when exiting the compiler_directive production.
	ExitCompiler_directive(c *Compiler_directiveContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitModule_declaration is called when exiting the module_declaration production.
	ExitModule_declaration(c *Module_declarationContext)

	// ExitModule_identifier is called when exiting the module_identifier production.
	ExitModule_identifier(c *Module_identifierContext)

	// ExitModule_interface is called when exiting the module_interface production.
	ExitModule_interface(c *Module_interfaceContext)

	// ExitModule_parameter_interface is called when exiting the module_parameter_interface production.
	ExitModule_parameter_interface(c *Module_parameter_interfaceContext)

	// ExitModule_port_interface is called when exiting the module_port_interface production.
	ExitModule_port_interface(c *Module_port_interfaceContext)

	// ExitModule_item_star is called when exiting the module_item_star production.
	ExitModule_item_star(c *Module_item_starContext)

	// ExitModule_item is called when exiting the module_item production.
	ExitModule_item(c *Module_itemContext)

	// ExitColon_module_identifier is called when exiting the colon_module_identifier production.
	ExitColon_module_identifier(c *Colon_module_identifierContext)

	// ExitPackage_declaration is called when exiting the package_declaration production.
	ExitPackage_declaration(c *Package_declarationContext)

	// ExitPackage_identifier is called when exiting the package_identifier production.
	ExitPackage_identifier(c *Package_identifierContext)

	// ExitColon_package_identifier is called when exiting the colon_package_identifier production.
	ExitColon_package_identifier(c *Colon_package_identifierContext)

	// ExitPackage_item_star is called when exiting the package_item_star production.
	ExitPackage_item_star(c *Package_item_starContext)

	// ExitPackage_item is called when exiting the package_item production.
	ExitPackage_item(c *Package_itemContext)

	// ExitImport_package is called when exiting the import_package production.
	ExitImport_package(c *Import_packageContext)

	// ExitPackage_item_identifier is called when exiting the package_item_identifier production.
	ExitPackage_item_identifier(c *Package_item_identifierContext)

	// ExitParameter_item_semicolon is called when exiting the parameter_item_semicolon production.
	ExitParameter_item_semicolon(c *Parameter_item_semicolonContext)

	// ExitParameter_item is called when exiting the parameter_item production.
	ExitParameter_item(c *Parameter_itemContext)

	// ExitAttr_port_item_semicolon is called when exiting the attr_port_item_semicolon production.
	ExitAttr_port_item_semicolon(c *Attr_port_item_semicolonContext)

	// ExitAttr_variable_item_semicolon is called when exiting the attr_variable_item_semicolon production.
	ExitAttr_variable_item_semicolon(c *Attr_variable_item_semicolonContext)

	// ExitVariable_item is called when exiting the variable_item production.
	ExitVariable_item(c *Variable_itemContext)

	// ExitSubroutine_item_semicolon is called when exiting the subroutine_item_semicolon production.
	ExitSubroutine_item_semicolon(c *Subroutine_item_semicolonContext)

	// ExitSubroutine_item is called when exiting the subroutine_item production.
	ExitSubroutine_item(c *Subroutine_itemContext)

	// ExitAttr_construct_item is called when exiting the attr_construct_item production.
	ExitAttr_construct_item(c *Attr_construct_itemContext)

	// ExitConstruct_item is called when exiting the construct_item production.
	ExitConstruct_item(c *Construct_itemContext)

	// ExitAttr_component_item is called when exiting the attr_component_item production.
	ExitAttr_component_item(c *Attr_component_itemContext)

	// ExitComponent_item is called when exiting the component_item production.
	ExitComponent_item(c *Component_itemContext)

	// ExitCompiler_item is called when exiting the compiler_item production.
	ExitCompiler_item(c *Compiler_itemContext)

	// ExitType_item is called when exiting the type_item production.
	ExitType_item(c *Type_itemContext)

	// ExitNull_item is called when exiting the null_item production.
	ExitNull_item(c *Null_itemContext)

	// ExitList_of_interface_parameters is called when exiting the list_of_interface_parameters production.
	ExitList_of_interface_parameters(c *List_of_interface_parametersContext)

	// ExitList_of_parameter_declarations is called when exiting the list_of_parameter_declarations production.
	ExitList_of_parameter_declarations(c *List_of_parameter_declarationsContext)

	// ExitComma_parameter_declaration_star is called when exiting the comma_parameter_declaration_star production.
	ExitComma_parameter_declaration_star(c *Comma_parameter_declaration_starContext)

	// ExitComma_parameter_declaration is called when exiting the comma_parameter_declaration production.
	ExitComma_parameter_declaration(c *Comma_parameter_declarationContext)

	// ExitList_of_parameter_descriptions is called when exiting the list_of_parameter_descriptions production.
	ExitList_of_parameter_descriptions(c *List_of_parameter_descriptionsContext)

	// ExitParam_declaration is called when exiting the param_declaration production.
	ExitParam_declaration(c *Param_declarationContext)

	// ExitParam_description is called when exiting the param_description production.
	ExitParam_description(c *Param_descriptionContext)

	// ExitParameter_declaration is called when exiting the parameter_declaration production.
	ExitParameter_declaration(c *Parameter_declarationContext)

	// ExitLocal_parameter_declaration is called when exiting the local_parameter_declaration production.
	ExitLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// ExitParameter_override is called when exiting the parameter_override production.
	ExitParameter_override(c *Parameter_overrideContext)

	// ExitList_of_tf_interface_ports is called when exiting the list_of_tf_interface_ports production.
	ExitList_of_tf_interface_ports(c *List_of_tf_interface_portsContext)

	// ExitList_of_tf_port_declarations is called when exiting the list_of_tf_port_declarations production.
	ExitList_of_tf_port_declarations(c *List_of_tf_port_declarationsContext)

	// ExitList_of_tf_port_declarations_comma is called when exiting the list_of_tf_port_declarations_comma production.
	ExitList_of_tf_port_declarations_comma(c *List_of_tf_port_declarations_commaContext)

	// ExitComma_attr_tf_port_declaration_star is called when exiting the comma_attr_tf_port_declaration_star production.
	ExitComma_attr_tf_port_declaration_star(c *Comma_attr_tf_port_declaration_starContext)

	// ExitComma_attr_tf_port_declaration is called when exiting the comma_attr_tf_port_declaration production.
	ExitComma_attr_tf_port_declaration(c *Comma_attr_tf_port_declarationContext)

	// ExitList_of_tf_port_declarations_semicolon is called when exiting the list_of_tf_port_declarations_semicolon production.
	ExitList_of_tf_port_declarations_semicolon(c *List_of_tf_port_declarations_semicolonContext)

	// ExitAttr_tf_port_declaration_semicolon_plus is called when exiting the attr_tf_port_declaration_semicolon_plus production.
	ExitAttr_tf_port_declaration_semicolon_plus(c *Attr_tf_port_declaration_semicolon_plusContext)

	// ExitAttr_tf_port_declaration_semicolon_star is called when exiting the attr_tf_port_declaration_semicolon_star production.
	ExitAttr_tf_port_declaration_semicolon_star(c *Attr_tf_port_declaration_semicolon_starContext)

	// ExitAttr_tf_port_declaration_semicolon is called when exiting the attr_tf_port_declaration_semicolon production.
	ExitAttr_tf_port_declaration_semicolon(c *Attr_tf_port_declaration_semicolonContext)

	// ExitAttr_tf_port_declaration is called when exiting the attr_tf_port_declaration production.
	ExitAttr_tf_port_declaration(c *Attr_tf_port_declarationContext)

	// ExitTf_port_declaration is called when exiting the tf_port_declaration production.
	ExitTf_port_declaration(c *Tf_port_declarationContext)

	// ExitList_of_interface_ports is called when exiting the list_of_interface_ports production.
	ExitList_of_interface_ports(c *List_of_interface_portsContext)

	// ExitList_of_port_identifiers is called when exiting the list_of_port_identifiers production.
	ExitList_of_port_identifiers(c *List_of_port_identifiersContext)

	// ExitComma_port_identifier_star is called when exiting the comma_port_identifier_star production.
	ExitComma_port_identifier_star(c *Comma_port_identifier_starContext)

	// ExitComma_port_identifier is called when exiting the comma_port_identifier production.
	ExitComma_port_identifier(c *Comma_port_identifierContext)

	// ExitPort_identifier is called when exiting the port_identifier production.
	ExitPort_identifier(c *Port_identifierContext)

	// ExitList_of_port_declarations is called when exiting the list_of_port_declarations production.
	ExitList_of_port_declarations(c *List_of_port_declarationsContext)

	// ExitList_of_port_declarations_comma is called when exiting the list_of_port_declarations_comma production.
	ExitList_of_port_declarations_comma(c *List_of_port_declarations_commaContext)

	// ExitComma_attr_port_declaration_star is called when exiting the comma_attr_port_declaration_star production.
	ExitComma_attr_port_declaration_star(c *Comma_attr_port_declaration_starContext)

	// ExitComma_attr_port_declaration is called when exiting the comma_attr_port_declaration production.
	ExitComma_attr_port_declaration(c *Comma_attr_port_declarationContext)

	// ExitList_of_port_declarations_semicolon is called when exiting the list_of_port_declarations_semicolon production.
	ExitList_of_port_declarations_semicolon(c *List_of_port_declarations_semicolonContext)

	// ExitAttr_port_declaration_semicolon_plus is called when exiting the attr_port_declaration_semicolon_plus production.
	ExitAttr_port_declaration_semicolon_plus(c *Attr_port_declaration_semicolon_plusContext)

	// ExitAttr_port_declaration_semicolon_star is called when exiting the attr_port_declaration_semicolon_star production.
	ExitAttr_port_declaration_semicolon_star(c *Attr_port_declaration_semicolon_starContext)

	// ExitAttr_port_declaration_semicolon is called when exiting the attr_port_declaration_semicolon production.
	ExitAttr_port_declaration_semicolon(c *Attr_port_declaration_semicolonContext)

	// ExitAttr_port_declaration is called when exiting the attr_port_declaration production.
	ExitAttr_port_declaration(c *Attr_port_declarationContext)

	// ExitPort_declaration is called when exiting the port_declaration production.
	ExitPort_declaration(c *Port_declarationContext)

	// ExitPort_description is called when exiting the port_description production.
	ExitPort_description(c *Port_descriptionContext)

	// ExitInout_description is called when exiting the inout_description production.
	ExitInout_description(c *Inout_descriptionContext)

	// ExitInput_description is called when exiting the input_description production.
	ExitInput_description(c *Input_descriptionContext)

	// ExitOutput_description is called when exiting the output_description production.
	ExitOutput_description(c *Output_descriptionContext)

	// ExitRef_description is called when exiting the ref_description production.
	ExitRef_description(c *Ref_descriptionContext)

	// ExitTf_declaration is called when exiting the tf_declaration production.
	ExitTf_declaration(c *Tf_declarationContext)

	// ExitInout_declaration is called when exiting the inout_declaration production.
	ExitInout_declaration(c *Inout_declarationContext)

	// ExitInput_declaration is called when exiting the input_declaration production.
	ExitInput_declaration(c *Input_declarationContext)

	// ExitOutput_declaration is called when exiting the output_declaration production.
	ExitOutput_declaration(c *Output_declarationContext)

	// ExitRef_declaration is called when exiting the ref_declaration production.
	ExitRef_declaration(c *Ref_declarationContext)

	// ExitUser_type is called when exiting the user_type production.
	ExitUser_type(c *User_typeContext)

	// ExitUser_type_identifer is called when exiting the user_type_identifer production.
	ExitUser_type_identifer(c *User_type_identiferContext)

	// ExitDimension_plus is called when exiting the dimension_plus production.
	ExitDimension_plus(c *Dimension_plusContext)

	// ExitDimension_star is called when exiting the dimension_star production.
	ExitDimension_star(c *Dimension_starContext)

	// ExitDimension is called when exiting the dimension production.
	ExitDimension(c *DimensionContext)

	// ExitRange_expression is called when exiting the range_expression production.
	ExitRange_expression(c *Range_expressionContext)

	// ExitIndex_expression is called when exiting the index_expression production.
	ExitIndex_expression(c *Index_expressionContext)

	// ExitSb_range is called when exiting the sb_range production.
	ExitSb_range(c *Sb_rangeContext)

	// ExitBase_increment_range is called when exiting the base_increment_range production.
	ExitBase_increment_range(c *Base_increment_rangeContext)

	// ExitBase_decrement_range is called when exiting the base_decrement_range production.
	ExitBase_decrement_range(c *Base_decrement_rangeContext)

	// ExitBase_expression is called when exiting the base_expression production.
	ExitBase_expression(c *Base_expressionContext)

	// ExitNet_type is called when exiting the net_type production.
	ExitNet_type(c *Net_typeContext)

	// ExitDrive_strength is called when exiting the drive_strength production.
	ExitDrive_strength(c *Drive_strengthContext)

	// ExitDrive_strength_value_0 is called when exiting the drive_strength_value_0 production.
	ExitDrive_strength_value_0(c *Drive_strength_value_0Context)

	// ExitDrive_strength_value_1 is called when exiting the drive_strength_value_1 production.
	ExitDrive_strength_value_1(c *Drive_strength_value_1Context)

	// ExitStrength0 is called when exiting the strength0 production.
	ExitStrength0(c *Strength0Context)

	// ExitStrength1 is called when exiting the strength1 production.
	ExitStrength1(c *Strength1Context)

	// ExitHighz0 is called when exiting the highz0 production.
	ExitHighz0(c *Highz0Context)

	// ExitHighz1 is called when exiting the highz1 production.
	ExitHighz1(c *Highz1Context)

	// ExitCharge_strength is called when exiting the charge_strength production.
	ExitCharge_strength(c *Charge_strengthContext)

	// ExitCharge_size is called when exiting the charge_size production.
	ExitCharge_size(c *Charge_sizeContext)

	// ExitList_of_variable_descriptions is called when exiting the list_of_variable_descriptions production.
	ExitList_of_variable_descriptions(c *List_of_variable_descriptionsContext)

	// ExitComma_variable_description_star is called when exiting the comma_variable_description_star production.
	ExitComma_variable_description_star(c *Comma_variable_description_starContext)

	// ExitComma_variable_description is called when exiting the comma_variable_description production.
	ExitComma_variable_description(c *Comma_variable_descriptionContext)

	// ExitVariable_description is called when exiting the variable_description production.
	ExitVariable_description(c *Variable_descriptionContext)

	// ExitVariable_identifier is called when exiting the variable_identifier production.
	ExitVariable_identifier(c *Variable_identifierContext)

	// ExitList_of_hierarchical_variable_descriptions is called when exiting the list_of_hierarchical_variable_descriptions production.
	ExitList_of_hierarchical_variable_descriptions(c *List_of_hierarchical_variable_descriptionsContext)

	// ExitComma_hierarchical_variable_description_star is called when exiting the comma_hierarchical_variable_description_star production.
	ExitComma_hierarchical_variable_description_star(c *Comma_hierarchical_variable_description_starContext)

	// ExitComma_hierarchical_variable_description is called when exiting the comma_hierarchical_variable_description production.
	ExitComma_hierarchical_variable_description(c *Comma_hierarchical_variable_descriptionContext)

	// ExitHierarchical_variable_description is called when exiting the hierarchical_variable_description production.
	ExitHierarchical_variable_description(c *Hierarchical_variable_descriptionContext)

	// ExitHierarchical_variable_identifier is called when exiting the hierarchical_variable_identifier production.
	ExitHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// ExitNet_declaration is called when exiting the net_declaration production.
	ExitNet_declaration(c *Net_declarationContext)

	// ExitReg_declaration is called when exiting the reg_declaration production.
	ExitReg_declaration(c *Reg_declarationContext)

	// ExitLogic_declaration is called when exiting the logic_declaration production.
	ExitLogic_declaration(c *Logic_declarationContext)

	// ExitBits_type is called when exiting the bits_type production.
	ExitBits_type(c *Bits_typeContext)

	// ExitBits_declaration is called when exiting the bits_declaration production.
	ExitBits_declaration(c *Bits_declarationContext)

	// ExitInteger_declaration is called when exiting the integer_declaration production.
	ExitInteger_declaration(c *Integer_declarationContext)

	// ExitInt_declaration is called when exiting the int_declaration production.
	ExitInt_declaration(c *Int_declarationContext)

	// ExitReal_declaration is called when exiting the real_declaration production.
	ExitReal_declaration(c *Real_declarationContext)

	// ExitTime_declaration is called when exiting the time_declaration production.
	ExitTime_declaration(c *Time_declarationContext)

	// ExitRealtime_declaration is called when exiting the realtime_declaration production.
	ExitRealtime_declaration(c *Realtime_declarationContext)

	// ExitEvent_declaration is called when exiting the event_declaration production.
	ExitEvent_declaration(c *Event_declarationContext)

	// ExitGenvar_declaration is called when exiting the genvar_declaration production.
	ExitGenvar_declaration(c *Genvar_declarationContext)

	// ExitUsertype_variable_declaration is called when exiting the usertype_variable_declaration production.
	ExitUsertype_variable_declaration(c *Usertype_variable_declarationContext)

	// ExitString_declaration is called when exiting the string_declaration production.
	ExitString_declaration(c *String_declarationContext)

	// ExitStruct_declaration is called when exiting the struct_declaration production.
	ExitStruct_declaration(c *Struct_declarationContext)

	// ExitEnum_declaration is called when exiting the enum_declaration production.
	ExitEnum_declaration(c *Enum_declarationContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitFunction_type is called when exiting the function_type production.
	ExitFunction_type(c *Function_typeContext)

	// ExitFunction_identifier is called when exiting the function_identifier production.
	ExitFunction_identifier(c *Function_identifierContext)

	// ExitFunction_interface is called when exiting the function_interface production.
	ExitFunction_interface(c *Function_interfaceContext)

	// ExitFunction_item_declaration_star is called when exiting the function_item_declaration_star production.
	ExitFunction_item_declaration_star(c *Function_item_declaration_starContext)

	// ExitFunction_item_declaration_semicolon is called when exiting the function_item_declaration_semicolon production.
	ExitFunction_item_declaration_semicolon(c *Function_item_declaration_semicolonContext)

	// ExitFunction_item_declaration is called when exiting the function_item_declaration production.
	ExitFunction_item_declaration(c *Function_item_declarationContext)

	// ExitFunction_statement is called when exiting the function_statement production.
	ExitFunction_statement(c *Function_statementContext)

	// ExitColon_function_identifier is called when exiting the colon_function_identifier production.
	ExitColon_function_identifier(c *Colon_function_identifierContext)

	// ExitTask_declaration is called when exiting the task_declaration production.
	ExitTask_declaration(c *Task_declarationContext)

	// ExitTask_identifier is called when exiting the task_identifier production.
	ExitTask_identifier(c *Task_identifierContext)

	// ExitTask_interface is called when exiting the task_interface production.
	ExitTask_interface(c *Task_interfaceContext)

	// ExitTask_item_declaration_semicolon is called when exiting the task_item_declaration_semicolon production.
	ExitTask_item_declaration_semicolon(c *Task_item_declaration_semicolonContext)

	// ExitTask_item_declaration is called when exiting the task_item_declaration production.
	ExitTask_item_declaration(c *Task_item_declarationContext)

	// ExitTask_item_declaration_star is called when exiting the task_item_declaration_star production.
	ExitTask_item_declaration_star(c *Task_item_declaration_starContext)

	// ExitTask_statement is called when exiting the task_statement production.
	ExitTask_statement(c *Task_statementContext)

	// ExitStruct_item_semicolon is called when exiting the struct_item_semicolon production.
	ExitStruct_item_semicolon(c *Struct_item_semicolonContext)

	// ExitStruct_item_star is called when exiting the struct_item_star production.
	ExitStruct_item_star(c *Struct_item_starContext)

	// ExitStruct_item is called when exiting the struct_item production.
	ExitStruct_item(c *Struct_itemContext)

	// ExitStruct_type is called when exiting the struct_type production.
	ExitStruct_type(c *Struct_typeContext)

	// ExitEnum_type is called when exiting the enum_type production.
	ExitEnum_type(c *Enum_typeContext)

	// ExitList_of_enum_items is called when exiting the list_of_enum_items production.
	ExitList_of_enum_items(c *List_of_enum_itemsContext)

	// ExitEnum_item is called when exiting the enum_item production.
	ExitEnum_item(c *Enum_itemContext)

	// ExitEnum_identifier is called when exiting the enum_identifier production.
	ExitEnum_identifier(c *Enum_identifierContext)

	// ExitComma_enum_item_star is called when exiting the comma_enum_item_star production.
	ExitComma_enum_item_star(c *Comma_enum_item_starContext)

	// ExitComma_enum_item is called when exiting the comma_enum_item production.
	ExitComma_enum_item(c *Comma_enum_itemContext)

	// ExitEnumerated_type is called when exiting the enumerated_type production.
	ExitEnumerated_type(c *Enumerated_typeContext)

	// ExitModule_instantiation is called when exiting the module_instantiation production.
	ExitModule_instantiation(c *Module_instantiationContext)

	// ExitParameter_interface_assignments is called when exiting the parameter_interface_assignments production.
	ExitParameter_interface_assignments(c *Parameter_interface_assignmentsContext)

	// ExitList_of_interface_assignments is called when exiting the list_of_interface_assignments production.
	ExitList_of_interface_assignments(c *List_of_interface_assignmentsContext)

	// ExitList_of_ordered_interface_assignments is called when exiting the list_of_ordered_interface_assignments production.
	ExitList_of_ordered_interface_assignments(c *List_of_ordered_interface_assignmentsContext)

	// ExitComma_ordered_interface_assignment_star is called when exiting the comma_ordered_interface_assignment_star production.
	ExitComma_ordered_interface_assignment_star(c *Comma_ordered_interface_assignment_starContext)

	// ExitComma_ordered_interface_assignment is called when exiting the comma_ordered_interface_assignment production.
	ExitComma_ordered_interface_assignment(c *Comma_ordered_interface_assignmentContext)

	// ExitOrdered_interface_assignment is called when exiting the ordered_interface_assignment production.
	ExitOrdered_interface_assignment(c *Ordered_interface_assignmentContext)

	// ExitList_of_named_interface_assignments is called when exiting the list_of_named_interface_assignments production.
	ExitList_of_named_interface_assignments(c *List_of_named_interface_assignmentsContext)

	// ExitComma_named_interface_assignment_star is called when exiting the comma_named_interface_assignment_star production.
	ExitComma_named_interface_assignment_star(c *Comma_named_interface_assignment_starContext)

	// ExitComma_named_interface_assignment is called when exiting the comma_named_interface_assignment production.
	ExitComma_named_interface_assignment(c *Comma_named_interface_assignmentContext)

	// ExitNamed_interface_assignment is called when exiting the named_interface_assignment production.
	ExitNamed_interface_assignment(c *Named_interface_assignmentContext)

	// ExitList_of_module_instances is called when exiting the list_of_module_instances production.
	ExitList_of_module_instances(c *List_of_module_instancesContext)

	// ExitComma_module_instance_star is called when exiting the comma_module_instance_star production.
	ExitComma_module_instance_star(c *Comma_module_instance_starContext)

	// ExitComma_module_instance is called when exiting the comma_module_instance production.
	ExitComma_module_instance(c *Comma_module_instanceContext)

	// ExitModule_instance is called when exiting the module_instance production.
	ExitModule_instance(c *Module_instanceContext)

	// ExitModule_instance_identifier is called when exiting the module_instance_identifier production.
	ExitModule_instance_identifier(c *Module_instance_identifierContext)

	// ExitArrayed_identifier is called when exiting the arrayed_identifier production.
	ExitArrayed_identifier(c *Arrayed_identifierContext)

	// ExitSimple_arrayed_identifier is called when exiting the simple_arrayed_identifier production.
	ExitSimple_arrayed_identifier(c *Simple_arrayed_identifierContext)

	// ExitEscaped_arrayed_identifier is called when exiting the escaped_arrayed_identifier production.
	ExitEscaped_arrayed_identifier(c *Escaped_arrayed_identifierContext)

	// ExitPort_interface_assignments is called when exiting the port_interface_assignments production.
	ExitPort_interface_assignments(c *Port_interface_assignmentsContext)

	// ExitDelay is called when exiting the delay production.
	ExitDelay(c *DelayContext)

	// ExitList_of_delay_values is called when exiting the list_of_delay_values production.
	ExitList_of_delay_values(c *List_of_delay_valuesContext)

	// ExitComma_delay_value_star is called when exiting the comma_delay_value_star production.
	ExitComma_delay_value_star(c *Comma_delay_value_starContext)

	// ExitComma_delay_value is called when exiting the comma_delay_value production.
	ExitComma_delay_value(c *Comma_delay_valueContext)

	// ExitDelay_value is called when exiting the delay_value production.
	ExitDelay_value(c *Delay_valueContext)

	// ExitPulldown_strength is called when exiting the pulldown_strength production.
	ExitPulldown_strength(c *Pulldown_strengthContext)

	// ExitPullup_strength is called when exiting the pullup_strength production.
	ExitPullup_strength(c *Pullup_strengthContext)

	// ExitGate_instance_identifier is called when exiting the gate_instance_identifier production.
	ExitGate_instance_identifier(c *Gate_instance_identifierContext)

	// ExitGate_instantiation is called when exiting the gate_instantiation production.
	ExitGate_instantiation(c *Gate_instantiationContext)

	// ExitEnable_gatetype is called when exiting the enable_gatetype production.
	ExitEnable_gatetype(c *Enable_gatetypeContext)

	// ExitMos_switchtype is called when exiting the mos_switchtype production.
	ExitMos_switchtype(c *Mos_switchtypeContext)

	// ExitCmos_switchtype is called when exiting the cmos_switchtype production.
	ExitCmos_switchtype(c *Cmos_switchtypeContext)

	// ExitN_output_gatetype is called when exiting the n_output_gatetype production.
	ExitN_output_gatetype(c *N_output_gatetypeContext)

	// ExitN_input_gatetype is called when exiting the n_input_gatetype production.
	ExitN_input_gatetype(c *N_input_gatetypeContext)

	// ExitPass_switchtype is called when exiting the pass_switchtype production.
	ExitPass_switchtype(c *Pass_switchtypeContext)

	// ExitPass_enable_switchtype is called when exiting the pass_enable_switchtype production.
	ExitPass_enable_switchtype(c *Pass_enable_switchtypeContext)

	// ExitPulldown_instantiation is called when exiting the pulldown_instantiation production.
	ExitPulldown_instantiation(c *Pulldown_instantiationContext)

	// ExitPullup_instantiation is called when exiting the pullup_instantiation production.
	ExitPullup_instantiation(c *Pullup_instantiationContext)

	// ExitEnable_instantiation is called when exiting the enable_instantiation production.
	ExitEnable_instantiation(c *Enable_instantiationContext)

	// ExitMos_instantiation is called when exiting the mos_instantiation production.
	ExitMos_instantiation(c *Mos_instantiationContext)

	// ExitCmos_instantiation is called when exiting the cmos_instantiation production.
	ExitCmos_instantiation(c *Cmos_instantiationContext)

	// ExitN_output_instantiation is called when exiting the n_output_instantiation production.
	ExitN_output_instantiation(c *N_output_instantiationContext)

	// ExitN_input_instantiation is called when exiting the n_input_instantiation production.
	ExitN_input_instantiation(c *N_input_instantiationContext)

	// ExitPass_instantiation is called when exiting the pass_instantiation production.
	ExitPass_instantiation(c *Pass_instantiationContext)

	// ExitPass_enable_instantiation is called when exiting the pass_enable_instantiation production.
	ExitPass_enable_instantiation(c *Pass_enable_instantiationContext)

	// ExitList_of_pull_gate_instance is called when exiting the list_of_pull_gate_instance production.
	ExitList_of_pull_gate_instance(c *List_of_pull_gate_instanceContext)

	// ExitList_of_enable_gate_instance is called when exiting the list_of_enable_gate_instance production.
	ExitList_of_enable_gate_instance(c *List_of_enable_gate_instanceContext)

	// ExitList_of_mos_switch_instance is called when exiting the list_of_mos_switch_instance production.
	ExitList_of_mos_switch_instance(c *List_of_mos_switch_instanceContext)

	// ExitList_of_cmos_switch_instance is called when exiting the list_of_cmos_switch_instance production.
	ExitList_of_cmos_switch_instance(c *List_of_cmos_switch_instanceContext)

	// ExitList_of_n_input_gate_instance is called when exiting the list_of_n_input_gate_instance production.
	ExitList_of_n_input_gate_instance(c *List_of_n_input_gate_instanceContext)

	// ExitList_of_n_output_gate_instance is called when exiting the list_of_n_output_gate_instance production.
	ExitList_of_n_output_gate_instance(c *List_of_n_output_gate_instanceContext)

	// ExitList_of_pass_switch_instance is called when exiting the list_of_pass_switch_instance production.
	ExitList_of_pass_switch_instance(c *List_of_pass_switch_instanceContext)

	// ExitList_of_pass_enable_switch_instance is called when exiting the list_of_pass_enable_switch_instance production.
	ExitList_of_pass_enable_switch_instance(c *List_of_pass_enable_switch_instanceContext)

	// ExitComma_pull_gate_instance_star is called when exiting the comma_pull_gate_instance_star production.
	ExitComma_pull_gate_instance_star(c *Comma_pull_gate_instance_starContext)

	// ExitComma_enable_gate_instance_star is called when exiting the comma_enable_gate_instance_star production.
	ExitComma_enable_gate_instance_star(c *Comma_enable_gate_instance_starContext)

	// ExitComma_mos_switch_instance_star is called when exiting the comma_mos_switch_instance_star production.
	ExitComma_mos_switch_instance_star(c *Comma_mos_switch_instance_starContext)

	// ExitComma_cmos_switch_instance_star is called when exiting the comma_cmos_switch_instance_star production.
	ExitComma_cmos_switch_instance_star(c *Comma_cmos_switch_instance_starContext)

	// ExitComma_n_input_gate_instance_star is called when exiting the comma_n_input_gate_instance_star production.
	ExitComma_n_input_gate_instance_star(c *Comma_n_input_gate_instance_starContext)

	// ExitComma_n_output_gate_instance_star is called when exiting the comma_n_output_gate_instance_star production.
	ExitComma_n_output_gate_instance_star(c *Comma_n_output_gate_instance_starContext)

	// ExitComma_pass_switch_instance_star is called when exiting the comma_pass_switch_instance_star production.
	ExitComma_pass_switch_instance_star(c *Comma_pass_switch_instance_starContext)

	// ExitComma_pass_enable_switch_instance_star is called when exiting the comma_pass_enable_switch_instance_star production.
	ExitComma_pass_enable_switch_instance_star(c *Comma_pass_enable_switch_instance_starContext)

	// ExitComma_pull_gate_instance is called when exiting the comma_pull_gate_instance production.
	ExitComma_pull_gate_instance(c *Comma_pull_gate_instanceContext)

	// ExitComma_enable_gate_instance is called when exiting the comma_enable_gate_instance production.
	ExitComma_enable_gate_instance(c *Comma_enable_gate_instanceContext)

	// ExitComma_mos_switch_instance is called when exiting the comma_mos_switch_instance production.
	ExitComma_mos_switch_instance(c *Comma_mos_switch_instanceContext)

	// ExitComma_cmos_switch_instance is called when exiting the comma_cmos_switch_instance production.
	ExitComma_cmos_switch_instance(c *Comma_cmos_switch_instanceContext)

	// ExitComma_n_input_gate_instance is called when exiting the comma_n_input_gate_instance production.
	ExitComma_n_input_gate_instance(c *Comma_n_input_gate_instanceContext)

	// ExitComma_n_output_gate_instance is called when exiting the comma_n_output_gate_instance production.
	ExitComma_n_output_gate_instance(c *Comma_n_output_gate_instanceContext)

	// ExitComma_pass_switch_instance is called when exiting the comma_pass_switch_instance production.
	ExitComma_pass_switch_instance(c *Comma_pass_switch_instanceContext)

	// ExitComma_pass_enable_switch_instance is called when exiting the comma_pass_enable_switch_instance production.
	ExitComma_pass_enable_switch_instance(c *Comma_pass_enable_switch_instanceContext)

	// ExitPull_gate_instance is called when exiting the pull_gate_instance production.
	ExitPull_gate_instance(c *Pull_gate_instanceContext)

	// ExitEnable_gate_instance is called when exiting the enable_gate_instance production.
	ExitEnable_gate_instance(c *Enable_gate_instanceContext)

	// ExitMos_switch_instance is called when exiting the mos_switch_instance production.
	ExitMos_switch_instance(c *Mos_switch_instanceContext)

	// ExitCmos_switch_instance is called when exiting the cmos_switch_instance production.
	ExitCmos_switch_instance(c *Cmos_switch_instanceContext)

	// ExitN_input_gate_instance is called when exiting the n_input_gate_instance production.
	ExitN_input_gate_instance(c *N_input_gate_instanceContext)

	// ExitN_output_gate_instance is called when exiting the n_output_gate_instance production.
	ExitN_output_gate_instance(c *N_output_gate_instanceContext)

	// ExitPass_switch_instance is called when exiting the pass_switch_instance production.
	ExitPass_switch_instance(c *Pass_switch_instanceContext)

	// ExitPass_enable_switch_instance is called when exiting the pass_enable_switch_instance production.
	ExitPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// ExitPull_gate_interface is called when exiting the pull_gate_interface production.
	ExitPull_gate_interface(c *Pull_gate_interfaceContext)

	// ExitEnable_gate_interface is called when exiting the enable_gate_interface production.
	ExitEnable_gate_interface(c *Enable_gate_interfaceContext)

	// ExitMos_switch_interface is called when exiting the mos_switch_interface production.
	ExitMos_switch_interface(c *Mos_switch_interfaceContext)

	// ExitCmos_switch_interface is called when exiting the cmos_switch_interface production.
	ExitCmos_switch_interface(c *Cmos_switch_interfaceContext)

	// ExitN_input_gate_interface is called when exiting the n_input_gate_interface production.
	ExitN_input_gate_interface(c *N_input_gate_interfaceContext)

	// ExitN_output_gate_interface is called when exiting the n_output_gate_interface production.
	ExitN_output_gate_interface(c *N_output_gate_interfaceContext)

	// ExitPass_switch_interface is called when exiting the pass_switch_interface production.
	ExitPass_switch_interface(c *Pass_switch_interfaceContext)

	// ExitPass_enable_switch_interface is called when exiting the pass_enable_switch_interface production.
	ExitPass_enable_switch_interface(c *Pass_enable_switch_interfaceContext)

	// ExitList_of_input_terminals is called when exiting the list_of_input_terminals production.
	ExitList_of_input_terminals(c *List_of_input_terminalsContext)

	// ExitList_of_output_terminals is called when exiting the list_of_output_terminals production.
	ExitList_of_output_terminals(c *List_of_output_terminalsContext)

	// ExitComma_input_terminal_star is called when exiting the comma_input_terminal_star production.
	ExitComma_input_terminal_star(c *Comma_input_terminal_starContext)

	// ExitComma_output_terminal_star is called when exiting the comma_output_terminal_star production.
	ExitComma_output_terminal_star(c *Comma_output_terminal_starContext)

	// ExitComma_input_terminal is called when exiting the comma_input_terminal production.
	ExitComma_input_terminal(c *Comma_input_terminalContext)

	// ExitComma_output_terminal is called when exiting the comma_output_terminal production.
	ExitComma_output_terminal(c *Comma_output_terminalContext)

	// ExitEnable_terminal is called when exiting the enable_terminal production.
	ExitEnable_terminal(c *Enable_terminalContext)

	// ExitInput_terminal is called when exiting the input_terminal production.
	ExitInput_terminal(c *Input_terminalContext)

	// ExitInout_terminal is called when exiting the inout_terminal production.
	ExitInout_terminal(c *Inout_terminalContext)

	// ExitNcontrol_terminal is called when exiting the ncontrol_terminal production.
	ExitNcontrol_terminal(c *Ncontrol_terminalContext)

	// ExitOutput_terminal is called when exiting the output_terminal production.
	ExitOutput_terminal(c *Output_terminalContext)

	// ExitPcontrol_terminal is called when exiting the pcontrol_terminal production.
	ExitPcontrol_terminal(c *Pcontrol_terminalContext)

	// ExitStatement_star is called when exiting the statement_star production.
	ExitStatement_star(c *Statement_starContext)

	// ExitStatement_semicolon is called when exiting the statement_semicolon production.
	ExitStatement_semicolon(c *Statement_semicolonContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment_statement is called when exiting the assignment_statement production.
	ExitAssignment_statement(c *Assignment_statementContext)

	// ExitFlow_control_statement is called when exiting the flow_control_statement production.
	ExitFlow_control_statement(c *Flow_control_statementContext)

	// ExitBlock_statement is called when exiting the block_statement production.
	ExitBlock_statement(c *Block_statementContext)

	// ExitTask_call_statement is called when exiting the task_call_statement production.
	ExitTask_call_statement(c *Task_call_statementContext)

	// ExitEvent_statement is called when exiting the event_statement production.
	ExitEvent_statement(c *Event_statementContext)

	// ExitProcedural_statement is called when exiting the procedural_statement production.
	ExitProcedural_statement(c *Procedural_statementContext)

	// ExitExpression_statement is called when exiting the expression_statement production.
	ExitExpression_statement(c *Expression_statementContext)

	// ExitSubroutine_statement is called when exiting the subroutine_statement production.
	ExitSubroutine_statement(c *Subroutine_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitNull_statement is called when exiting the null_statement production.
	ExitNull_statement(c *Null_statementContext)

	// ExitProcedural_continuous_assignments is called when exiting the procedural_continuous_assignments production.
	ExitProcedural_continuous_assignments(c *Procedural_continuous_assignmentsContext)

	// ExitAssign_statement is called when exiting the assign_statement production.
	ExitAssign_statement(c *Assign_statementContext)

	// ExitDeassign_statement is called when exiting the deassign_statement production.
	ExitDeassign_statement(c *Deassign_statementContext)

	// ExitForce_statement is called when exiting the force_statement production.
	ExitForce_statement(c *Force_statementContext)

	// ExitRelease_statement is called when exiting the release_statement production.
	ExitRelease_statement(c *Release_statementContext)

	// ExitProcedural_timing_control_statement is called when exiting the procedural_timing_control_statement production.
	ExitProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// ExitProperty_statement is called when exiting the property_statement production.
	ExitProperty_statement(c *Property_statementContext)

	// ExitDisable_condition_statement is called when exiting the disable_condition_statement production.
	ExitDisable_condition_statement(c *Disable_condition_statementContext)

	// ExitProperty_expression is called when exiting the property_expression production.
	ExitProperty_expression(c *Property_expressionContext)

	// ExitProcedural_assertion_statement is called when exiting the procedural_assertion_statement production.
	ExitProcedural_assertion_statement(c *Procedural_assertion_statementContext)

	// ExitAssert_else_statement is called when exiting the assert_else_statement production.
	ExitAssert_else_statement(c *Assert_else_statementContext)

	// ExitAssert_statement is called when exiting the assert_statement production.
	ExitAssert_statement(c *Assert_statementContext)

	// ExitSystem_task_enable is called when exiting the system_task_enable production.
	ExitSystem_task_enable(c *System_task_enableContext)

	// ExitSystem_task_identifier is called when exiting the system_task_identifier production.
	ExitSystem_task_identifier(c *System_task_identifierContext)

	// ExitTask_interface_assignments is called when exiting the task_interface_assignments production.
	ExitTask_interface_assignments(c *Task_interface_assignmentsContext)

	// ExitTask_enable is called when exiting the task_enable production.
	ExitTask_enable(c *Task_enableContext)

	// ExitHierarchical_task_identifier is called when exiting the hierarchical_task_identifier production.
	ExitHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// ExitDisable_statement is called when exiting the disable_statement production.
	ExitDisable_statement(c *Disable_statementContext)

	// ExitHierarchical_block_identifier is called when exiting the hierarchical_block_identifier production.
	ExitHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// ExitVariable_lvalue is called when exiting the variable_lvalue production.
	ExitVariable_lvalue(c *Variable_lvalueContext)

	// ExitHierarchical_variable_lvalue is called when exiting the hierarchical_variable_lvalue production.
	ExitHierarchical_variable_lvalue(c *Hierarchical_variable_lvalueContext)

	// ExitVariable_concatenation is called when exiting the variable_concatenation production.
	ExitVariable_concatenation(c *Variable_concatenationContext)

	// ExitVariable_concatenation_value is called when exiting the variable_concatenation_value production.
	ExitVariable_concatenation_value(c *Variable_concatenation_valueContext)

	// ExitComma_vcv_star is called when exiting the comma_vcv_star production.
	ExitComma_vcv_star(c *Comma_vcv_starContext)

	// ExitBlocking_assignment is called when exiting the blocking_assignment production.
	ExitBlocking_assignment(c *Blocking_assignmentContext)

	// ExitNonblocking_assignment is called when exiting the nonblocking_assignment production.
	ExitNonblocking_assignment(c *Nonblocking_assignmentContext)

	// ExitPrefix_assignment is called when exiting the prefix_assignment production.
	ExitPrefix_assignment(c *Prefix_assignmentContext)

	// ExitPostfix_assignment is called when exiting the postfix_assignment production.
	ExitPostfix_assignment(c *Postfix_assignmentContext)

	// ExitOperator_assignment is called when exiting the operator_assignment production.
	ExitOperator_assignment(c *Operator_assignmentContext)

	// ExitDeclarative_assignment is called when exiting the declarative_assignment production.
	ExitDeclarative_assignment(c *Declarative_assignmentContext)

	// ExitDelay_or_event_control is called when exiting the delay_or_event_control production.
	ExitDelay_or_event_control(c *Delay_or_event_controlContext)

	// ExitDelay_control is called when exiting the delay_control production.
	ExitDelay_control(c *Delay_controlContext)

	// ExitEvent_control is called when exiting the event_control production.
	ExitEvent_control(c *Event_controlContext)

	// ExitEvent_control_identifier is called when exiting the event_control_identifier production.
	ExitEvent_control_identifier(c *Event_control_identifierContext)

	// ExitEvent_control_expression is called when exiting the event_control_expression production.
	ExitEvent_control_expression(c *Event_control_expressionContext)

	// ExitEvent_expression is called when exiting the event_expression production.
	ExitEvent_expression(c *Event_expressionContext)

	// ExitSingle_event_expression is called when exiting the single_event_expression production.
	ExitSingle_event_expression(c *Single_event_expressionContext)

	// ExitEvent_expression_edgespec is called when exiting the event_expression_edgespec production.
	ExitEvent_expression_edgespec(c *Event_expression_edgespecContext)

	// ExitEvent_expression_or is called when exiting the event_expression_or production.
	ExitEvent_expression_or(c *Event_expression_orContext)

	// ExitList_of_event_expression_comma is called when exiting the list_of_event_expression_comma production.
	ExitList_of_event_expression_comma(c *List_of_event_expression_commaContext)

	// ExitComma_event_expression_star is called when exiting the comma_event_expression_star production.
	ExitComma_event_expression_star(c *Comma_event_expression_starContext)

	// ExitComma_event_expression is called when exiting the comma_event_expression production.
	ExitComma_event_expression(c *Comma_event_expressionContext)

	// ExitList_of_event_expression_or is called when exiting the list_of_event_expression_or production.
	ExitList_of_event_expression_or(c *List_of_event_expression_orContext)

	// ExitOr_event_expression_star is called when exiting the or_event_expression_star production.
	ExitOr_event_expression_star(c *Or_event_expression_starContext)

	// ExitOr_event_expression is called when exiting the or_event_expression production.
	ExitOr_event_expression(c *Or_event_expressionContext)

	// ExitEvent_control_wildcard is called when exiting the event_control_wildcard production.
	ExitEvent_control_wildcard(c *Event_control_wildcardContext)

	// ExitRepeat_event_control is called when exiting the repeat_event_control production.
	ExitRepeat_event_control(c *Repeat_event_controlContext)

	// ExitEvent_trigger is called when exiting the event_trigger production.
	ExitEvent_trigger(c *Event_triggerContext)

	// ExitHierarchical_event_identifier is called when exiting the hierarchical_event_identifier production.
	ExitHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// ExitEvent_identifier is called when exiting the event_identifier production.
	ExitEvent_identifier(c *Event_identifierContext)

	// ExitWait_statement is called when exiting the wait_statement production.
	ExitWait_statement(c *Wait_statementContext)

	// ExitAttr_generated_instantiation is called when exiting the attr_generated_instantiation production.
	ExitAttr_generated_instantiation(c *Attr_generated_instantiationContext)

	// ExitGenerated_instantiation is called when exiting the generated_instantiation production.
	ExitGenerated_instantiation(c *Generated_instantiationContext)

	// ExitGenerate_item_star is called when exiting the generate_item_star production.
	ExitGenerate_item_star(c *Generate_item_starContext)

	// ExitGenerate_item is called when exiting the generate_item production.
	ExitGenerate_item(c *Generate_itemContext)

	// ExitGenerate_block is called when exiting the generate_block production.
	ExitGenerate_block(c *Generate_blockContext)

	// ExitGenerate_colon_block_identifier0 is called when exiting the generate_colon_block_identifier0 production.
	ExitGenerate_colon_block_identifier0(c *Generate_colon_block_identifier0Context)

	// ExitGenerate_colon_block_identifier1 is called when exiting the generate_colon_block_identifier1 production.
	ExitGenerate_colon_block_identifier1(c *Generate_colon_block_identifier1Context)

	// ExitGenerate_colon_block_identifier is called when exiting the generate_colon_block_identifier production.
	ExitGenerate_colon_block_identifier(c *Generate_colon_block_identifierContext)

	// ExitGenerate_block_identifier is called when exiting the generate_block_identifier production.
	ExitGenerate_block_identifier(c *Generate_block_identifierContext)

	// ExitGenerate_conditional_statement is called when exiting the generate_conditional_statement production.
	ExitGenerate_conditional_statement(c *Generate_conditional_statementContext)

	// ExitGenerate_if_statement is called when exiting the generate_if_statement production.
	ExitGenerate_if_statement(c *Generate_if_statementContext)

	// ExitGenerate_else_statement is called when exiting the generate_else_statement production.
	ExitGenerate_else_statement(c *Generate_else_statementContext)

	// ExitGenerate_loop_statement is called when exiting the generate_loop_statement production.
	ExitGenerate_loop_statement(c *Generate_loop_statementContext)

	// ExitGenerate_forever_loop_statement is called when exiting the generate_forever_loop_statement production.
	ExitGenerate_forever_loop_statement(c *Generate_forever_loop_statementContext)

	// ExitGenerate_repeat_loop_statement is called when exiting the generate_repeat_loop_statement production.
	ExitGenerate_repeat_loop_statement(c *Generate_repeat_loop_statementContext)

	// ExitGenerate_while_loop_statement is called when exiting the generate_while_loop_statement production.
	ExitGenerate_while_loop_statement(c *Generate_while_loop_statementContext)

	// ExitGenerate_do_loop_statement is called when exiting the generate_do_loop_statement production.
	ExitGenerate_do_loop_statement(c *Generate_do_loop_statementContext)

	// ExitGenerate_for_loop_statement is called when exiting the generate_for_loop_statement production.
	ExitGenerate_for_loop_statement(c *Generate_for_loop_statementContext)

	// ExitGenerate_case_statement is called when exiting the generate_case_statement production.
	ExitGenerate_case_statement(c *Generate_case_statementContext)

	// ExitGenerate_case_item_star is called when exiting the generate_case_item_star production.
	ExitGenerate_case_item_star(c *Generate_case_item_starContext)

	// ExitGenerate_case_item is called when exiting the generate_case_item production.
	ExitGenerate_case_item(c *Generate_case_itemContext)

	// ExitConditional_statement is called when exiting the conditional_statement production.
	ExitConditional_statement(c *Conditional_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitElse_statement is called when exiting the else_statement production.
	ExitElse_statement(c *Else_statementContext)

	// ExitConditional_expression is called when exiting the conditional_expression production.
	ExitConditional_expression(c *Conditional_expressionContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitForever_loop_statement is called when exiting the forever_loop_statement production.
	ExitForever_loop_statement(c *Forever_loop_statementContext)

	// ExitRepeat_loop_statement is called when exiting the repeat_loop_statement production.
	ExitRepeat_loop_statement(c *Repeat_loop_statementContext)

	// ExitWhile_loop_statement is called when exiting the while_loop_statement production.
	ExitWhile_loop_statement(c *While_loop_statementContext)

	// ExitDo_loop_statement is called when exiting the do_loop_statement production.
	ExitDo_loop_statement(c *Do_loop_statementContext)

	// ExitFor_loop_statement is called when exiting the for_loop_statement production.
	ExitFor_loop_statement(c *For_loop_statementContext)

	// ExitLoop_init_assignment is called when exiting the loop_init_assignment production.
	ExitLoop_init_assignment(c *Loop_init_assignmentContext)

	// ExitLoop_terminate_expression is called when exiting the loop_terminate_expression production.
	ExitLoop_terminate_expression(c *Loop_terminate_expressionContext)

	// ExitLoop_step_assignment is called when exiting the loop_step_assignment production.
	ExitLoop_step_assignment(c *Loop_step_assignmentContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitCase_item_star is called when exiting the case_item_star production.
	ExitCase_item_star(c *Case_item_starContext)

	// ExitCase_item is called when exiting the case_item production.
	ExitCase_item(c *Case_itemContext)

	// ExitCase_switch is called when exiting the case_switch production.
	ExitCase_switch(c *Case_switchContext)

	// ExitCase_item_key is called when exiting the case_item_key production.
	ExitCase_item_key(c *Case_item_keyContext)

	// ExitCase_item_key_expression is called when exiting the case_item_key_expression production.
	ExitCase_item_key_expression(c *Case_item_key_expressionContext)

	// ExitComma_case_item_key_expression is called when exiting the comma_case_item_key_expression production.
	ExitComma_case_item_key_expression(c *Comma_case_item_key_expressionContext)

	// ExitComma_case_item_key_expression_star is called when exiting the comma_case_item_key_expression_star production.
	ExitComma_case_item_key_expression_star(c *Comma_case_item_key_expression_starContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSingle_expression is called when exiting the single_expression production.
	ExitSingle_expression(c *Single_expressionContext)

	// ExitPrimary_range is called when exiting the primary_range production.
	ExitPrimary_range(c *Primary_rangeContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitUnary_expression is called when exiting the unary_expression production.
	ExitUnary_expression(c *Unary_expressionContext)

	// ExitUnary_post_assign_expression is called when exiting the unary_post_assign_expression production.
	ExitUnary_post_assign_expression(c *Unary_post_assign_expressionContext)

	// ExitUnary_pre_assign_expression is called when exiting the unary_pre_assign_expression production.
	ExitUnary_pre_assign_expression(c *Unary_pre_assign_expressionContext)

	// ExitBinary_expression is called when exiting the binary_expression production.
	ExitBinary_expression(c *Binary_expressionContext)

	// ExitTernary_expression is called when exiting the ternary_expression production.
	ExitTernary_expression(c *Ternary_expressionContext)

	// ExitMintypmax_expression is called when exiting the mintypmax_expression production.
	ExitMintypmax_expression(c *Mintypmax_expressionContext)

	// ExitStructured_value is called when exiting the structured_value production.
	ExitStructured_value(c *Structured_valueContext)

	// ExitArrayed_structured_value is called when exiting the arrayed_structured_value production.
	ExitArrayed_structured_value(c *Arrayed_structured_valueContext)

	// ExitArrayed_structure_item is called when exiting the arrayed_structure_item production.
	ExitArrayed_structure_item(c *Arrayed_structure_itemContext)

	// ExitComma_arrayed_structure_item is called when exiting the comma_arrayed_structure_item production.
	ExitComma_arrayed_structure_item(c *Comma_arrayed_structure_itemContext)

	// ExitComma_arrayed_structure_item_star is called when exiting the comma_arrayed_structure_item_star production.
	ExitComma_arrayed_structure_item_star(c *Comma_arrayed_structure_item_starContext)

	// ExitArrayed_structure_item_plus is called when exiting the arrayed_structure_item_plus production.
	ExitArrayed_structure_item_plus(c *Arrayed_structure_item_plusContext)

	// ExitVariable_type_cast is called when exiting the variable_type_cast production.
	ExitVariable_type_cast(c *Variable_type_castContext)

	// ExitWidth_type_cast is called when exiting the width_type_cast production.
	ExitWidth_type_cast(c *Width_type_castContext)

	// ExitSign_type_cast is called when exiting the sign_type_cast production.
	ExitSign_type_cast(c *Sign_type_castContext)

	// ExitNull_type_cast is called when exiting the null_type_cast production.
	ExitNull_type_cast(c *Null_type_castContext)

	// ExitVariable_type is called when exiting the variable_type production.
	ExitVariable_type(c *Variable_typeContext)

	// ExitType_cast_identifier is called when exiting the type_cast_identifier production.
	ExitType_cast_identifier(c *Type_cast_identifierContext)

	// ExitType_cast_expression is called when exiting the type_cast_expression production.
	ExitType_cast_expression(c *Type_cast_expressionContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitHierarchical_function_identifier is called when exiting the hierarchical_function_identifier production.
	ExitHierarchical_function_identifier(c *Hierarchical_function_identifierContext)

	// ExitFunction_interface_assignments is called when exiting the function_interface_assignments production.
	ExitFunction_interface_assignments(c *Function_interface_assignmentsContext)

	// ExitSystem_function_call is called when exiting the system_function_call production.
	ExitSystem_function_call(c *System_function_callContext)

	// ExitSystem_function_identifier is called when exiting the system_function_identifier production.
	ExitSystem_function_identifier(c *System_function_identifierContext)

	// ExitConstant_function_call is called when exiting the constant_function_call production.
	ExitConstant_function_call(c *Constant_function_callContext)

	// ExitImported_function_call is called when exiting the imported_function_call production.
	ExitImported_function_call(c *Imported_function_callContext)

	// ExitImported_function_hierarchical_identifier is called when exiting the imported_function_hierarchical_identifier production.
	ExitImported_function_hierarchical_identifier(c *Imported_function_hierarchical_identifierContext)

	// ExitPrimary_hierarchical_identifier is called when exiting the primary_hierarchical_identifier production.
	ExitPrimary_hierarchical_identifier(c *Primary_hierarchical_identifierContext)

	// ExitPrimary_imported_hierarchical_identifier is called when exiting the primary_imported_hierarchical_identifier production.
	ExitPrimary_imported_hierarchical_identifier(c *Primary_imported_hierarchical_identifierContext)

	// ExitImported_hierarchical_identifier is called when exiting the imported_hierarchical_identifier production.
	ExitImported_hierarchical_identifier(c *Imported_hierarchical_identifierContext)

	// ExitParenthesis_expression is called when exiting the parenthesis_expression production.
	ExitParenthesis_expression(c *Parenthesis_expressionContext)

	// ExitConcatenation is called when exiting the concatenation production.
	ExitConcatenation(c *ConcatenationContext)

	// ExitMultiple_concatenation is called when exiting the multiple_concatenation production.
	ExitMultiple_concatenation(c *Multiple_concatenationContext)

	// ExitComma_expression_plus is called when exiting the comma_expression_plus production.
	ExitComma_expression_plus(c *Comma_expression_plusContext)

	// ExitComma_expression_star is called when exiting the comma_expression_star production.
	ExitComma_expression_star(c *Comma_expression_starContext)

	// ExitTypedef_declaration is called when exiting the typedef_declaration production.
	ExitTypedef_declaration(c *Typedef_declarationContext)

	// ExitTypedef_identifier is called when exiting the typedef_identifier production.
	ExitTypedef_identifier(c *Typedef_identifierContext)

	// ExitTypedef_definition is called when exiting the typedef_definition production.
	ExitTypedef_definition(c *Typedef_definitionContext)

	// ExitTypedef_definition_type is called when exiting the typedef_definition_type production.
	ExitTypedef_definition_type(c *Typedef_definition_typeContext)

	// ExitComplex_type is called when exiting the complex_type production.
	ExitComplex_type(c *Complex_typeContext)

	// ExitTypedef_type is called when exiting the typedef_type production.
	ExitTypedef_type(c *Typedef_typeContext)

	// ExitPar_block is called when exiting the par_block production.
	ExitPar_block(c *Par_blockContext)

	// ExitSeq_block is called when exiting the seq_block production.
	ExitSeq_block(c *Seq_blockContext)

	// ExitBlock_identifier is called when exiting the block_identifier production.
	ExitBlock_identifier(c *Block_identifierContext)

	// ExitColon_block_identifier is called when exiting the colon_block_identifier production.
	ExitColon_block_identifier(c *Colon_block_identifierContext)

	// ExitBlock_item_declaration_star is called when exiting the block_item_declaration_star production.
	ExitBlock_item_declaration_star(c *Block_item_declaration_starContext)

	// ExitBlock_item_declaration_semicolon is called when exiting the block_item_declaration_semicolon production.
	ExitBlock_item_declaration_semicolon(c *Block_item_declaration_semicolonContext)

	// ExitBlock_item_declaration is called when exiting the block_item_declaration production.
	ExitBlock_item_declaration(c *Block_item_declarationContext)

	// ExitJoin_keyword is called when exiting the join_keyword production.
	ExitJoin_keyword(c *Join_keywordContext)

	// ExitContinuous_assign is called when exiting the continuous_assign production.
	ExitContinuous_assign(c *Continuous_assignContext)

	// ExitList_of_variable_assignments is called when exiting the list_of_variable_assignments production.
	ExitList_of_variable_assignments(c *List_of_variable_assignmentsContext)

	// ExitComma_variable_assignment_star is called when exiting the comma_variable_assignment_star production.
	ExitComma_variable_assignment_star(c *Comma_variable_assignment_starContext)

	// ExitComma_variable_assignment is called when exiting the comma_variable_assignment production.
	ExitComma_variable_assignment(c *Comma_variable_assignmentContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitInitial_construct is called when exiting the initial_construct production.
	ExitInitial_construct(c *Initial_constructContext)

	// ExitFinal_construct is called when exiting the final_construct production.
	ExitFinal_construct(c *Final_constructContext)

	// ExitAlways_keyword is called when exiting the always_keyword production.
	ExitAlways_keyword(c *Always_keywordContext)

	// ExitAlways_construct is called when exiting the always_construct production.
	ExitAlways_construct(c *Always_constructContext)

	// ExitAttribute_instance_star is called when exiting the attribute_instance_star production.
	ExitAttribute_instance_star(c *Attribute_instance_starContext)

	// ExitAttribute_instance is called when exiting the attribute_instance production.
	ExitAttribute_instance(c *Attribute_instanceContext)

	// ExitAttr_spec_star is called when exiting the attr_spec_star production.
	ExitAttr_spec_star(c *Attr_spec_starContext)

	// ExitAttr_spec is called when exiting the attr_spec production.
	ExitAttr_spec(c *Attr_specContext)

	// ExitAttr_name is called when exiting the attr_name production.
	ExitAttr_name(c *Attr_nameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitHierarchical_identifier is called when exiting the hierarchical_identifier production.
	ExitHierarchical_identifier(c *Hierarchical_identifierContext)

	// ExitDot_hierarchical_identifier_branch_item_star is called when exiting the dot_hierarchical_identifier_branch_item_star production.
	ExitDot_hierarchical_identifier_branch_item_star(c *Dot_hierarchical_identifier_branch_item_starContext)

	// ExitDot_hierarchical_identifier_branch_item is called when exiting the dot_hierarchical_identifier_branch_item production.
	ExitDot_hierarchical_identifier_branch_item(c *Dot_hierarchical_identifier_branch_itemContext)

	// ExitHierarchical_identifier_branch_item is called when exiting the hierarchical_identifier_branch_item production.
	ExitHierarchical_identifier_branch_item(c *Hierarchical_identifier_branch_itemContext)

	// ExitTimescale_compiler_directive is called when exiting the timescale_compiler_directive production.
	ExitTimescale_compiler_directive(c *Timescale_compiler_directiveContext)

	// ExitTimeunit_directive is called when exiting the timeunit_directive production.
	ExitTimeunit_directive(c *Timeunit_directiveContext)

	// ExitTimeprecision_directive is called when exiting the timeprecision_directive production.
	ExitTimeprecision_directive(c *Timeprecision_directiveContext)

	// ExitDefault_nettype_statement is called when exiting the default_nettype_statement production.
	ExitDefault_nettype_statement(c *Default_nettype_statementContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitIntegral_number is called when exiting the integral_number production.
	ExitIntegral_number(c *Integral_numberContext)

	// ExitReal_number is called when exiting the real_number production.
	ExitReal_number(c *Real_numberContext)
}
