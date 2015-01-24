
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 34 : // ['"','"']
				return 2
			case r == 39 : // [''',''']
				return 3
			case r == 40 : // ['(','(']
				return 4
			case r == 41 : // [')',')']
				return 5
			case r == 44 : // [',',',']
				return 6
			case r == 45 : // ['-','-']
				return 7
			case r == 46 : // ['.','.']
				return 8
			case r == 47 : // ['/','/']
				return 9
			case r == 48 : // ['0','0']
				return 10
			case 49 <= r && r <= 57 : // ['1','9']
				return 11
			case r == 61 : // ['=','=']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 95 : // ['_','_']
				return 14
			case r == 96 : // ['`','`']
				return 15
			case 97 <= r && r <= 104 : // ['a','h']
				return 16
			case r == 105 : // ['i','i']
				return 17
			case 106 <= r && r <= 107 : // ['j','k']
				return 16
			case r == 108 : // ['l','l']
				return 18
			case 109 <= r && r <= 111 : // ['m','o']
				return 16
			case r == 112 : // ['p','p']
				return 19
			case 113 <= r && r <= 122 : // ['q','z']
				return 16
			case r == 123 : // ['{','{']
				return 20
			case r == 125 : // ['}','}']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 22
			case r == 92 : // ['\','\']
				return 23
			
			
			default:
				return 24
			}
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			}
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 48 : // ['0','0']
				return 27
			case 49 <= r && r <= 57 : // ['1','9']
				return 28
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 29
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 30
			case r == 47 : // ['/','/']
				return 31
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 32
			case 48 <= r && r <= 55 : // ['0','7']
				return 33
			case 56 <= r && r <= 57 : // ['8','9']
				return 34
			case r == 69 : // ['E','E']
				return 35
			case r == 88 : // ['X','X']
				return 36
			case r == 101 : // ['e','e']
				return 35
			case r == 120 : // ['x','x']
				return 36
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 32
			case 48 <= r && r <= 57 : // ['0','9']
				return 11
			case r == 69 : // ['E','E']
				return 35
			case r == 101 : // ['e','e']
				return 35
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 96 : // ['`','`']
				return 41
			
			
			default:
				return 15
			}
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 108 : // ['a','l']
				return 40
			case r == 109 : // ['m','m']
				return 42
			case 110 <= r && r <= 122 : // ['n','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 100 : // ['a','d']
				return 40
			case r == 101 : // ['e','e']
				return 43
			case 102 <= r && r <= 122 : // ['f','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 113 : // ['a','q']
				return 40
			case r == 114 : // ['r','r']
				return 44
			case 115 <= r && r <= 122 : // ['s','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 45
			case r == 39 : // [''',''']
				return 45
			case 48 <= r && r <= 55 : // ['0','7']
				return 46
			case r == 85 : // ['U','U']
				return 47
			case r == 92 : // ['\','\']
				return 45
			case r == 97 : // ['a','a']
				return 45
			case r == 98 : // ['b','b']
				return 45
			case r == 102 : // ['f','f']
				return 45
			case r == 110 : // ['n','n']
				return 45
			case r == 114 : // ['r','r']
				return 45
			case r == 116 : // ['t','t']
				return 45
			case r == 117 : // ['u','u']
				return 48
			case r == 118 : // ['v','v']
				return 45
			case r == 120 : // ['x','x']
				return 49
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 22
			case r == 92 : // ['\','\']
				return 23
			
			
			default:
				return 24
			}
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 50
			case r == 39 : // [''',''']
				return 50
			case 48 <= r && r <= 55 : // ['0','7']
				return 51
			case r == 85 : // ['U','U']
				return 52
			case r == 92 : // ['\','\']
				return 50
			case r == 97 : // ['a','a']
				return 50
			case r == 98 : // ['b','b']
				return 50
			case r == 102 : // ['f','f']
				return 50
			case r == 110 : // ['n','n']
				return 50
			case r == 114 : // ['r','r']
				return 50
			case r == 116 : // ['t','t']
				return 50
			case r == 117 : // ['u','u']
				return 53
			case r == 118 : // ['v','v']
				return 50
			case r == 120 : // ['x','x']
				return 54
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 55
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 56
			case r == 88 : // ['X','X']
				return 57
			case r == 120 : // ['x','x']
				return 57
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 58
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 29
			case r == 69 : // ['E','E']
				return 59
			case r == 101 : // ['e','e']
				return 59
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 60
			
			
			default:
				return 30
			}
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 61
			
			
			default:
				return 31
			}
			
		},
	
		// S32
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 62
			case r == 69 : // ['E','E']
				return 63
			case r == 101 : // ['e','e']
				return 63
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 32
			case 48 <= r && r <= 55 : // ['0','7']
				return 33
			case 56 <= r && r <= 57 : // ['8','9']
				return 34
			case r == 69 : // ['E','E']
				return 35
			case r == 101 : // ['e','e']
				return 35
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 32
			case 48 <= r && r <= 57 : // ['0','9']
				return 34
			case r == 69 : // ['E','E']
				return 35
			case r == 101 : // ['e','e']
				return 35
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case r == 43 : // ['+','+']
				return 64
			case r == 45 : // ['-','-']
				return 64
			case 48 <= r && r <= 57 : // ['0','9']
				return 65
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 66
			case 65 <= r && r <= 70 : // ['A','F']
				return 66
			case 97 <= r && r <= 102 : // ['a','f']
				return 66
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 111 : // ['a','o']
				return 40
			case r == 112 : // ['p','p']
				return 67
			case 113 <= r && r <= 122 : // ['q','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 115 : // ['a','s']
				return 40
			case r == 116 : // ['t','t']
				return 68
			case 117 <= r && r <= 122 : // ['u','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 104 : // ['a','h']
				return 40
			case r == 105 : // ['i','i']
				return 69
			case 106 <= r && r <= 122 : // ['j','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 22
			case r == 92 : // ['\','\']
				return 23
			
			
			default:
				return 24
			}
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 70
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 71
			case 65 <= r && r <= 70 : // ['A','F']
				return 71
			case 97 <= r && r <= 102 : // ['a','f']
				return 71
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 72
			case 65 <= r && r <= 70 : // ['A','F']
				return 72
			case 97 <= r && r <= 102 : // ['a','f']
				return 72
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 73
			case 65 <= r && r <= 70 : // ['A','F']
				return 73
			case 97 <= r && r <= 102 : // ['a','f']
				return 73
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 55
			
			
			
			}
			return NoState
			
		},
	
		// S51
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 74
			
			
			
			}
			return NoState
			
		},
	
		// S52
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 75
			case 65 <= r && r <= 70 : // ['A','F']
				return 75
			case 97 <= r && r <= 102 : // ['a','f']
				return 75
			
			
			
			}
			return NoState
			
		},
	
		// S53
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 76
			case 65 <= r && r <= 70 : // ['A','F']
				return 76
			case 97 <= r && r <= 102 : // ['a','f']
				return 76
			
			
			
			}
			return NoState
			
		},
	
		// S54
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 77
			case 65 <= r && r <= 70 : // ['A','F']
				return 77
			case 97 <= r && r <= 102 : // ['a','f']
				return 77
			
			
			
			}
			return NoState
			
		},
	
		// S55
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S56
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 56
			
			
			
			}
			return NoState
			
		},
	
		// S57
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 78
			case 65 <= r && r <= 70 : // ['A','F']
				return 78
			case 97 <= r && r <= 102 : // ['a','f']
				return 78
			
			
			
			}
			return NoState
			
		},
	
		// S58
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 58
			
			
			
			}
			return NoState
			
		},
	
		// S59
		func(r rune) int {
			switch {
			case r == 43 : // ['+','+']
				return 79
			case r == 45 : // ['-','-']
				return 79
			case 48 <= r && r <= 57 : // ['0','9']
				return 80
			
			
			
			}
			return NoState
			
		},
	
		// S60
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 60
			case r == 47 : // ['/','/']
				return 81
			
			
			default:
				return 30
			}
			
		},
	
		// S61
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S62
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 62
			case r == 69 : // ['E','E']
				return 82
			case r == 101 : // ['e','e']
				return 82
			
			
			
			}
			return NoState
			
		},
	
		// S63
		func(r rune) int {
			switch {
			case r == 43 : // ['+','+']
				return 83
			case r == 45 : // ['-','-']
				return 83
			case 48 <= r && r <= 57 : // ['0','9']
				return 84
			
			
			
			}
			return NoState
			
		},
	
		// S64
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 65
			
			
			
			}
			return NoState
			
		},
	
		// S65
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 65
			
			
			
			}
			return NoState
			
		},
	
		// S66
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 66
			case 65 <= r && r <= 70 : // ['A','F']
				return 66
			case 97 <= r && r <= 102 : // ['a','f']
				return 66
			
			
			
			}
			return NoState
			
		},
	
		// S67
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 110 : // ['a','n']
				return 40
			case r == 111 : // ['o','o']
				return 85
			case 112 <= r && r <= 122 : // ['p','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S68
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S69
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 109 : // ['a','m']
				return 40
			case r == 110 : // ['n','n']
				return 86
			case 111 <= r && r <= 122 : // ['o','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S70
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 87
			
			
			
			}
			return NoState
			
		},
	
		// S71
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 88
			case 65 <= r && r <= 70 : // ['A','F']
				return 88
			case 97 <= r && r <= 102 : // ['a','f']
				return 88
			
			
			
			}
			return NoState
			
		},
	
		// S72
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 89
			case 65 <= r && r <= 70 : // ['A','F']
				return 89
			case 97 <= r && r <= 102 : // ['a','f']
				return 89
			
			
			
			}
			return NoState
			
		},
	
		// S73
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 90
			case 65 <= r && r <= 70 : // ['A','F']
				return 90
			case 97 <= r && r <= 102 : // ['a','f']
				return 90
			
			
			
			}
			return NoState
			
		},
	
		// S74
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 91
			
			
			
			}
			return NoState
			
		},
	
		// S75
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 92
			case 65 <= r && r <= 70 : // ['A','F']
				return 92
			case 97 <= r && r <= 102 : // ['a','f']
				return 92
			
			
			
			}
			return NoState
			
		},
	
		// S76
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 93
			case 65 <= r && r <= 70 : // ['A','F']
				return 93
			case 97 <= r && r <= 102 : // ['a','f']
				return 93
			
			
			
			}
			return NoState
			
		},
	
		// S77
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 94
			case 65 <= r && r <= 70 : // ['A','F']
				return 94
			case 97 <= r && r <= 102 : // ['a','f']
				return 94
			
			
			
			}
			return NoState
			
		},
	
		// S78
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 78
			case 65 <= r && r <= 70 : // ['A','F']
				return 78
			case 97 <= r && r <= 102 : // ['a','f']
				return 78
			
			
			
			}
			return NoState
			
		},
	
		// S79
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 80
			
			
			
			}
			return NoState
			
		},
	
		// S80
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 80
			
			
			
			}
			return NoState
			
		},
	
		// S81
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S82
		func(r rune) int {
			switch {
			case r == 43 : // ['+','+']
				return 95
			case r == 45 : // ['-','-']
				return 95
			case 48 <= r && r <= 57 : // ['0','9']
				return 96
			
			
			
			}
			return NoState
			
		},
	
		// S83
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 84
			
			
			
			}
			return NoState
			
		},
	
		// S84
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 84
			
			
			
			}
			return NoState
			
		},
	
		// S85
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 113 : // ['a','q']
				return 40
			case r == 114 : // ['r','r']
				return 97
			case 115 <= r && r <= 122 : // ['s','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S86
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 115 : // ['a','s']
				return 40
			case r == 116 : // ['t','t']
				return 98
			case 117 <= r && r <= 122 : // ['u','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S87
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 22
			case r == 92 : // ['\','\']
				return 23
			
			
			default:
				return 24
			}
			
		},
	
		// S88
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 99
			case 65 <= r && r <= 70 : // ['A','F']
				return 99
			case 97 <= r && r <= 102 : // ['a','f']
				return 99
			
			
			
			}
			return NoState
			
		},
	
		// S89
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 100
			case 65 <= r && r <= 70 : // ['A','F']
				return 100
			case 97 <= r && r <= 102 : // ['a','f']
				return 100
			
			
			
			}
			return NoState
			
		},
	
		// S90
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 22
			case r == 92 : // ['\','\']
				return 23
			
			
			default:
				return 24
			}
			
		},
	
		// S91
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 55
			
			
			
			}
			return NoState
			
		},
	
		// S92
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 101
			case 65 <= r && r <= 70 : // ['A','F']
				return 101
			case 97 <= r && r <= 102 : // ['a','f']
				return 101
			
			
			
			}
			return NoState
			
		},
	
		// S93
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 102
			case 65 <= r && r <= 70 : // ['A','F']
				return 102
			case 97 <= r && r <= 102 : // ['a','f']
				return 102
			
			
			
			}
			return NoState
			
		},
	
		// S94
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 55
			
			
			
			}
			return NoState
			
		},
	
		// S95
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 96
			
			
			
			}
			return NoState
			
		},
	
		// S96
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 96
			
			
			
			}
			return NoState
			
		},
	
		// S97
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 115 : // ['a','s']
				return 40
			case r == 116 : // ['t','t']
				return 103
			case 117 <= r && r <= 122 : // ['u','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S98
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 107 : // ['a','k']
				return 40
			case r == 108 : // ['l','l']
				return 104
			case 109 <= r && r <= 122 : // ['m','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S99
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 105
			case 65 <= r && r <= 70 : // ['A','F']
				return 105
			case 97 <= r && r <= 102 : // ['a','f']
				return 105
			
			
			
			}
			return NoState
			
		},
	
		// S100
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 106
			case 65 <= r && r <= 70 : // ['A','F']
				return 106
			case 97 <= r && r <= 102 : // ['a','f']
				return 106
			
			
			
			}
			return NoState
			
		},
	
		// S101
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 107
			case 65 <= r && r <= 70 : // ['A','F']
				return 107
			case 97 <= r && r <= 102 : // ['a','f']
				return 107
			
			
			
			}
			return NoState
			
		},
	
		// S102
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 108
			case 65 <= r && r <= 70 : // ['A','F']
				return 108
			case 97 <= r && r <= 102 : // ['a','f']
				return 108
			
			
			
			}
			return NoState
			
		},
	
		// S103
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S104
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 109 : // ['a','m']
				return 40
			case r == 110 : // ['n','n']
				return 109
			case 111 <= r && r <= 122 : // ['o','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S105
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 110
			case 65 <= r && r <= 70 : // ['A','F']
				return 110
			case 97 <= r && r <= 102 : // ['a','f']
				return 110
			
			
			
			}
			return NoState
			
		},
	
		// S106
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 22
			case r == 92 : // ['\','\']
				return 23
			
			
			default:
				return 24
			}
			
		},
	
		// S107
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 111
			case 65 <= r && r <= 70 : // ['A','F']
				return 111
			case 97 <= r && r <= 102 : // ['a','f']
				return 111
			
			
			
			}
			return NoState
			
		},
	
		// S108
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 55
			
			
			
			}
			return NoState
			
		},
	
		// S109
		func(r rune) int {
			switch {
			case r == 40 : // ['(','(']
				return 112
			case 48 <= r && r <= 57 : // ['0','9']
				return 37
			case 65 <= r && r <= 90 : // ['A','Z']
				return 38
			case r == 95 : // ['_','_']
				return 39
			case 97 <= r && r <= 122 : // ['a','z']
				return 40
			
			
			
			}
			return NoState
			
		},
	
		// S110
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 113
			case 65 <= r && r <= 70 : // ['A','F']
				return 113
			case 97 <= r && r <= 102 : // ['a','f']
				return 113
			
			
			
			}
			return NoState
			
		},
	
		// S111
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 114
			case 65 <= r && r <= 70 : // ['A','F']
				return 114
			case 97 <= r && r <= 102 : // ['a','f']
				return 114
			
			
			
			}
			return NoState
			
		},
	
		// S112
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S113
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 115
			case 65 <= r && r <= 70 : // ['A','F']
				return 115
			case 97 <= r && r <= 102 : // ['a','f']
				return 115
			
			
			
			}
			return NoState
			
		},
	
		// S114
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 116
			case 65 <= r && r <= 70 : // ['A','F']
				return 116
			case 97 <= r && r <= 102 : // ['a','f']
				return 116
			
			
			
			}
			return NoState
			
		},
	
		// S115
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 117
			case 65 <= r && r <= 70 : // ['A','F']
				return 117
			case 97 <= r && r <= 102 : // ['a','f']
				return 117
			
			
			
			}
			return NoState
			
		},
	
		// S116
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 118
			case 65 <= r && r <= 70 : // ['A','F']
				return 118
			case 97 <= r && r <= 102 : // ['a','f']
				return 118
			
			
			
			}
			return NoState
			
		},
	
		// S117
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 22
			case r == 92 : // ['\','\']
				return 23
			
			
			default:
				return 24
			}
			
		},
	
		// S118
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 55
			
			
			
			}
			return NoState
			
		},
	
}
