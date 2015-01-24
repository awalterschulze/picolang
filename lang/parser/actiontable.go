
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(7),		/* let */
			shift(8),		/* id */
			nil,		/* = */
			shift(9),		/* import */
			nil,		/* string_lit */
			shift(10),		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			shift(7),		/* let */
			shift(8),		/* id */
			nil,		/* = */
			shift(9),		/* import */
			nil,		/* string_lit */
			shift(10),		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Statements */
			reduce(2),		/* let, reduce: Statements */
			reduce(2),		/* id, reduce: Statements */
			nil,		/* = */
			reduce(2),		/* import, reduce: Statements */
			nil,		/* string_lit */
			reduce(2),		/* println(, reduce: Statements */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Statement */
			reduce(3),		/* let, reduce: Statement */
			reduce(3),		/* id, reduce: Statement */
			nil,		/* = */
			reduce(3),		/* import, reduce: Statement */
			nil,		/* string_lit */
			reduce(3),		/* println(, reduce: Statement */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Statement */
			reduce(4),		/* let, reduce: Statement */
			reduce(4),		/* id, reduce: Statement */
			nil,		/* = */
			reduce(4),		/* import, reduce: Statement */
			nil,		/* string_lit */
			reduce(4),		/* println(, reduce: Statement */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Statement */
			reduce(5),		/* let, reduce: Statement */
			reduce(5),		/* id, reduce: Statement */
			nil,		/* = */
			reduce(5),		/* import, reduce: Statement */
			nil,		/* string_lit */
			reduce(5),		/* println(, reduce: Statement */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Statement */
			reduce(6),		/* let, reduce: Statement */
			reduce(6),		/* id, reduce: Statement */
			nil,		/* = */
			reduce(6),		/* import, reduce: Statement */
			nil,		/* string_lit */
			reduce(6),		/* println(, reduce: Statement */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(12),		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Expr */
			reduce(11),		/* let, reduce: Expr */
			reduce(11),		/* id, reduce: Expr */
			nil,		/* = */
			reduce(11),		/* import, reduce: Expr */
			nil,		/* string_lit */
			reduce(11),		/* println(, reduce: Expr */
			nil,		/* ) */
			shift(13),		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(14),		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(16),		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Statements */
			reduce(1),		/* let, reduce: Statements */
			reduce(1),		/* id, reduce: Statements */
			nil,		/* = */
			reduce(1),		/* import, reduce: Statements */
			nil,		/* string_lit */
			reduce(1),		/* println(, reduce: Statements */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			shift(17),		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* import */
			shift(19),		/* string_lit */
			nil,		/* println( */
			shift(20),		/* ) */
			nil,		/* ( */
			nil,		/* , */
			shift(23),		/* float_lit */
			shift(24),		/* singed_int_lit */
			shift(25),		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			shift(26),		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			shift(27),		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(11),		/* ), reduce: Expr */
			shift(28),		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(8),		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(15),		/* ), reduce: Param */
			nil,		/* ( */
			reduce(15),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(16),		/* ), reduce: Param */
			nil,		/* ( */
			reduce(16),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Expr */
			reduce(10),		/* let, reduce: Expr */
			reduce(10),		/* id, reduce: Expr */
			nil,		/* = */
			reduce(10),		/* import, reduce: Expr */
			nil,		/* string_lit */
			reduce(10),		/* println(, reduce: Expr */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			shift(30),		/* ) */
			nil,		/* ( */
			shift(31),		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(14),		/* ), reduce: Params */
			nil,		/* ( */
			reduce(14),		/* ,, reduce: Params */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(17),		/* ), reduce: Param */
			nil,		/* ( */
			reduce(17),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(18),		/* ), reduce: Param */
			nil,		/* ( */
			reduce(18),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(32),		/* id */
			nil,		/* = */
			nil,		/* import */
			shift(33),		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			shift(36),		/* float_lit */
			shift(37),		/* singed_int_lit */
			shift(38),		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Import */
			reduce(8),		/* let, reduce: Import */
			reduce(8),		/* id, reduce: Import */
			nil,		/* = */
			reduce(8),		/* import, reduce: Import */
			nil,		/* string_lit */
			reduce(8),		/* println(, reduce: Import */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Println */
			reduce(9),		/* let, reduce: Println */
			reduce(9),		/* id, reduce: Println */
			nil,		/* = */
			reduce(9),		/* import, reduce: Println */
			nil,		/* string_lit */
			reduce(9),		/* println(, reduce: Println */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* import */
			shift(19),		/* string_lit */
			nil,		/* println( */
			shift(39),		/* ) */
			nil,		/* ( */
			nil,		/* , */
			shift(23),		/* float_lit */
			shift(24),		/* singed_int_lit */
			shift(25),		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Let */
			reduce(7),		/* let, reduce: Let */
			reduce(7),		/* id, reduce: Let */
			nil,		/* = */
			reduce(7),		/* import, reduce: Let */
			nil,		/* string_lit */
			reduce(7),		/* println(, reduce: Let */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: Expr */
			reduce(12),		/* let, reduce: Expr */
			reduce(12),		/* id, reduce: Expr */
			nil,		/* = */
			reduce(12),		/* import, reduce: Expr */
			nil,		/* string_lit */
			reduce(12),		/* println(, reduce: Expr */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(18),		/* id */
			nil,		/* = */
			nil,		/* import */
			shift(19),		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			shift(23),		/* float_lit */
			shift(24),		/* singed_int_lit */
			shift(25),		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			reduce(15),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			reduce(15),		/* }, reduce: Param */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			reduce(16),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			reduce(16),		/* }, reduce: Param */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			shift(42),		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			shift(43),		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			reduce(14),		/* ,, reduce: Params */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			reduce(14),		/* }, reduce: Params */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			reduce(17),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			reduce(17),		/* }, reduce: Param */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			reduce(18),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			reduce(18),		/* }, reduce: Param */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(32),		/* id */
			nil,		/* = */
			nil,		/* import */
			shift(33),		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			shift(36),		/* float_lit */
			shift(37),		/* singed_int_lit */
			shift(38),		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(10),		/* ), reduce: Expr */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			shift(45),		/* ) */
			nil,		/* ( */
			shift(31),		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(13),		/* ), reduce: Params */
			nil,		/* ( */
			reduce(13),		/* ,, reduce: Params */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			shift(32),		/* id */
			nil,		/* = */
			nil,		/* import */
			shift(33),		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			nil,		/* , */
			shift(36),		/* float_lit */
			shift(37),		/* singed_int_lit */
			shift(38),		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(19),		/* ), reduce: Param */
			nil,		/* ( */
			reduce(19),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			shift(42),		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			shift(47),		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			reduce(12),		/* ), reduce: Expr */
			nil,		/* ( */
			nil,		/* , */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			nil,		/* } */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			reduce(13),		/* ,, reduce: Params */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			reduce(13),		/* }, reduce: Params */
			nil,		/* byte_elem */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* let */
			nil,		/* id */
			nil,		/* = */
			nil,		/* import */
			nil,		/* string_lit */
			nil,		/* println( */
			nil,		/* ) */
			nil,		/* ( */
			reduce(19),		/* ,, reduce: Param */
			nil,		/* float_lit */
			nil,		/* singed_int_lit */
			nil,		/* { */
			reduce(19),		/* }, reduce: Param */
			nil,		/* byte_elem */
			
		},

	},
	
}

