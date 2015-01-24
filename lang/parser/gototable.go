
/*
*/
package parser

const numNTSymbols = 9
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Statements
		2, // Statement
		3, // Let
		4, // Import
		5, // Println
		6, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Statements
		11, // Statement
		3, // Let
		4, // Import
		5, // Println
		6, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		15, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		21, // Params
		22, // Param
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		29, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		34, // Params
		35, // Param
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		40, // Params
		22, // Param
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		41, // Param
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		44, // Params
		35, // Param
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S42
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		46, // Param
		

	},
	gotoRow{ // S43
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S44
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S45
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S46
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	gotoRow{ // S47
		
		-1, // S'
		-1, // Statements
		-1, // Statement
		-1, // Let
		-1, // Import
		-1, // Println
		-1, // Expr
		-1, // Params
		-1, // Param
		

	},
	
}
