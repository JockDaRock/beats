//line parser.rl:1

// Code generated by ragel DO NOT EDIT.
package syslog

//line parser.go:8
const syslog_start int = 0
const syslog_first_final int = 1
const syslog_error int = -1

const syslog_en_main int = 0

//line parser.rl:9

// syslog
//<34>Oct 11 22:14:15 wopr su: 'su root' failed for foobar
//<13>Feb  5 17:32:18 10.0.0.99 Use the quad dmg.
func Parse(data []byte, event *event) {
	var p, cs int
	pe := len(data)
	tok := 0
	eof := len(data)

//line parser.go:28
	{
		cs = syslog_start
	}

//line parser.go:33
	{
		if (p) == (pe) {
			goto _test_eof
		}
		switch cs {
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		}
		goto st_out
	st_case_0:
		switch data[(p)] {
		case 60:
			goto tr1
		case 65:
			goto tr2
		case 70:
			goto tr3
		case 74:
			goto tr4
		case 77:
			goto tr5
		case 78:
			goto tr6
		case 79:
			goto tr7
		case 83:
			goto tr8
		case 101:
			goto tr9
		}
		goto tr0
	tr0:
//line parser.rl:20
		tok = p

		goto st1
	st1:
		if (p)++; (p) == (pe) {
			goto _test_eof1
		}
	st_case_1:
//line parser.go:228
		goto st1
	tr1:
//line parser.rl:20
		tok = p

		goto st2
	st2:
		if (p)++; (p) == (pe) {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:241
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto tr11
		}
		goto st1
	tr11:
//line parser.rl:20
		tok = p

		goto st3
	st3:
		if (p)++; (p) == (pe) {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:257
		if data[(p)] == 62 {
			goto tr13
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st4
		}
		goto st1
	st4:
		if (p)++; (p) == (pe) {
			goto _test_eof4
		}
	st_case_4:
		if data[(p)] == 62 {
			goto tr13
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st5
		}
		goto st1
	st5:
		if (p)++; (p) == (pe) {
			goto _test_eof5
		}
	st_case_5:
		if data[(p)] == 62 {
			goto tr13
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st6
		}
		goto st1
	st6:
		if (p)++; (p) == (pe) {
			goto _test_eof6
		}
	st_case_6:
		if data[(p)] == 62 {
			goto tr13
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st7
		}
		goto st1
	st7:
		if (p)++; (p) == (pe) {
			goto _test_eof7
		}
	st_case_7:
		if data[(p)] == 62 {
			goto tr13
		}
		goto st1
	tr13:
//line parser.rl:24
		event.SetPriority(data[tok:p])

		goto st8
	st8:
		if (p)++; (p) == (pe) {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:321
		switch data[(p)] {
		case 65:
			goto tr2
		case 70:
			goto tr3
		case 74:
			goto tr4
		case 77:
			goto tr5
		case 78:
			goto tr6
		case 79:
			goto tr7
		case 83:
			goto tr8
		case 101:
			goto tr9
		}
		goto tr0
	tr2:
//line parser.rl:20
		tok = p

		goto st9
	st9:
		if (p)++; (p) == (pe) {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:352
		switch data[(p)] {
		case 112:
			goto st10
		case 117:
			goto st41
		}
		goto st1
	st10:
		if (p)++; (p) == (pe) {
			goto _test_eof10
		}
	st_case_10:
		if data[(p)] == 114 {
			goto st11
		}
		goto st1
	st11:
		if (p)++; (p) == (pe) {
			goto _test_eof11
		}
	st_case_11:
		switch data[(p)] {
		case 32:
			goto tr20
		case 105:
			goto st39
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	tr20:
//line parser.rl:32
		event.SetMonth(data[tok:p])

		goto st12
	st12:
		if (p)++; (p) == (pe) {
			goto _test_eof12
		}
	st_case_12:
//line parser.go:395
		switch data[(p)] {
		case 32:
			goto st13
		case 51:
			goto tr24
		}
		switch {
		case data[(p)] < 49:
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto st13
			}
		case data[(p)] > 50:
			if 52 <= data[(p)] && data[(p)] <= 57 {
				goto tr25
			}
		default:
			goto tr23
		}
		goto st1
	st13:
		if (p)++; (p) == (pe) {
			goto _test_eof13
		}
	st_case_13:
		if 49 <= data[(p)] && data[(p)] <= 57 {
			goto tr25
		}
		goto st1
	tr25:
//line parser.rl:20
		tok = p

		goto st14
	st14:
		if (p)++; (p) == (pe) {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:435
		if data[(p)] == 32 {
			goto tr26
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr26
		}
		goto st1
	tr26:
//line parser.rl:36
		event.SetDay(data[tok:p])

		goto st15
	st15:
		if (p)++; (p) == (pe) {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:454
		if data[(p)] == 50 {
			goto tr28
		}
		if 48 <= data[(p)] && data[(p)] <= 49 {
			goto tr27
		}
		goto st1
	tr27:
//line parser.rl:20
		tok = p

		goto st16
	st16:
		if (p)++; (p) == (pe) {
			goto _test_eof16
		}
	st_case_16:
//line parser.go:473
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st17
		}
		goto st1
	st17:
		if (p)++; (p) == (pe) {
			goto _test_eof17
		}
	st_case_17:
		if data[(p)] == 58 {
			goto tr30
		}
		goto st1
	tr30:
//line parser.rl:40
		event.SetHour(data[tok:p])

		goto st18
	st18:
		if (p)++; (p) == (pe) {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:498
		if 48 <= data[(p)] && data[(p)] <= 53 {
			goto tr31
		}
		goto st1
	tr31:
//line parser.rl:20
		tok = p

		goto st19
	st19:
		if (p)++; (p) == (pe) {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:514
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st20
		}
		goto st1
	st20:
		if (p)++; (p) == (pe) {
			goto _test_eof20
		}
	st_case_20:
		if data[(p)] == 58 {
			goto tr33
		}
		goto st1
	tr33:
//line parser.rl:44
		event.SetMinute(data[tok:p])

		goto st21
	st21:
		if (p)++; (p) == (pe) {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:539
		if 48 <= data[(p)] && data[(p)] <= 53 {
			goto tr34
		}
		goto st1
	tr34:
//line parser.rl:20
		tok = p

		goto st22
	st22:
		if (p)++; (p) == (pe) {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:555
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st23
		}
		goto st1
	st23:
		if (p)++; (p) == (pe) {
			goto _test_eof23
		}
	st_case_23:
		switch data[(p)] {
		case 32:
			goto tr36
		case 46:
			goto tr37
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr36
		}
		goto st1
	tr36:
//line parser.rl:48
		event.SetSecond(data[tok:p])

		goto st24
	st24:
		if (p)++; (p) == (pe) {
			goto _test_eof24
		}
	st_case_24:
//line parser.go:586
		switch {
		case data[(p)] > 95:
			if 97 <= data[(p)] && data[(p)] <= 122 {
				goto tr38
			}
		case data[(p)] >= 46:
			goto tr38
		}
		goto tr0
	tr38:
//line parser.rl:20
		tok = p

		goto st25
	st25:
		if (p)++; (p) == (pe) {
			goto _test_eof25
		}
	st_case_25:
//line parser.go:607
		if data[(p)] == 32 {
			goto tr39
		}
		switch {
		case data[(p)] < 46:
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto tr39
			}
		case data[(p)] > 95:
			if 97 <= data[(p)] && data[(p)] <= 122 {
				goto st25
			}
		default:
			goto st25
		}
		goto st1
	tr39:
//line parser.rl:56
		event.SetHostname(data[tok:p])

		goto st26
	st26:
		if (p)++; (p) == (pe) {
			goto _test_eof26
		}
	st_case_26:
//line parser.go:635
		switch data[(p)] {
		case 32:
			goto tr0
		case 91:
			goto tr0
		case 93:
			goto tr0
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr0
		}
		goto tr41
	tr41:
//line parser.rl:20
		tok = p

		goto st27
	st27:
		if (p)++; (p) == (pe) {
			goto _test_eof27
		}
	st_case_27:
//line parser.go:659
		switch data[(p)] {
		case 32:
			goto st1
		case 58:
			goto tr43
		case 91:
			goto tr44
		case 93:
			goto st1
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto st1
		}
		goto st27
	tr43:
//line parser.rl:60
		event.SetProgram(data[tok:p])

		goto st28
	st28:
		if (p)++; (p) == (pe) {
			goto _test_eof28
		}
	st_case_28:
//line parser.go:685
		switch data[(p)] {
		case 32:
			goto st29
		case 58:
			goto tr43
		case 91:
			goto tr44
		case 93:
			goto st1
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto st29
		}
		goto st27
	st29:
		if (p)++; (p) == (pe) {
			goto _test_eof29
		}
	st_case_29:
		goto tr0
	tr44:
//line parser.rl:60
		event.SetProgram(data[tok:p])

		goto st30
	st30:
		if (p)++; (p) == (pe) {
			goto _test_eof30
		}
	st_case_30:
//line parser.go:717
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto tr46
		}
		goto st1
	tr46:
//line parser.rl:20
		tok = p

		goto st31
	st31:
		if (p)++; (p) == (pe) {
			goto _test_eof31
		}
	st_case_31:
//line parser.go:733
		if data[(p)] == 93 {
			goto tr48
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st31
		}
		goto st1
	tr48:
//line parser.rl:64
		event.SetPid(data[tok:p])

		goto st32
	st32:
		if (p)++; (p) == (pe) {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:752
		if data[(p)] == 58 {
			goto st33
		}
		goto st1
	st33:
		if (p)++; (p) == (pe) {
			goto _test_eof33
		}
	st_case_33:
		if data[(p)] == 32 {
			goto st29
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto st29
		}
		goto st1
	tr37:
//line parser.rl:48
		event.SetSecond(data[tok:p])

		goto st34
	st34:
		if (p)++; (p) == (pe) {
			goto _test_eof34
		}
	st_case_34:
//line parser.go:780
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st35
		}
		goto st1
	st35:
		if (p)++; (p) == (pe) {
			goto _test_eof35
		}
	st_case_35:
		if data[(p)] == 32 {
			goto st24
		}
		switch {
		case data[(p)] > 13:
			if 48 <= data[(p)] && data[(p)] <= 57 {
				goto st35
			}
		case data[(p)] >= 9:
			goto st24
		}
		goto st1
	tr28:
//line parser.rl:20
		tok = p

		goto st36
	st36:
		if (p)++; (p) == (pe) {
			goto _test_eof36
		}
	st_case_36:
//line parser.go:813
		if 48 <= data[(p)] && data[(p)] <= 51 {
			goto st17
		}
		goto st1
	tr23:
//line parser.rl:20
		tok = p

		goto st37
	st37:
		if (p)++; (p) == (pe) {
			goto _test_eof37
		}
	st_case_37:
//line parser.go:829
		if data[(p)] == 32 {
			goto tr26
		}
		switch {
		case data[(p)] > 13:
			if 48 <= data[(p)] && data[(p)] <= 57 {
				goto st14
			}
		case data[(p)] >= 9:
			goto tr26
		}
		goto st1
	tr24:
//line parser.rl:20
		tok = p

		goto st38
	st38:
		if (p)++; (p) == (pe) {
			goto _test_eof38
		}
	st_case_38:
//line parser.go:853
		if data[(p)] == 32 {
			goto tr26
		}
		switch {
		case data[(p)] > 13:
			if 48 <= data[(p)] && data[(p)] <= 49 {
				goto st14
			}
		case data[(p)] >= 9:
			goto tr26
		}
		goto st1
	st39:
		if (p)++; (p) == (pe) {
			goto _test_eof39
		}
	st_case_39:
		if data[(p)] == 108 {
			goto st40
		}
		goto st1
	st40:
		if (p)++; (p) == (pe) {
			goto _test_eof40
		}
	st_case_40:
		if data[(p)] == 32 {
			goto tr20
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st41:
		if (p)++; (p) == (pe) {
			goto _test_eof41
		}
	st_case_41:
		if data[(p)] == 103 {
			goto st42
		}
		goto st1
	st42:
		if (p)++; (p) == (pe) {
			goto _test_eof42
		}
	st_case_42:
		switch data[(p)] {
		case 32:
			goto tr20
		case 117:
			goto st43
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st43:
		if (p)++; (p) == (pe) {
			goto _test_eof43
		}
	st_case_43:
		if data[(p)] == 115 {
			goto st44
		}
		goto st1
	st44:
		if (p)++; (p) == (pe) {
			goto _test_eof44
		}
	st_case_44:
		if data[(p)] == 116 {
			goto st40
		}
		goto st1
	tr3:
//line parser.rl:20
		tok = p

		goto st45
	st45:
		if (p)++; (p) == (pe) {
			goto _test_eof45
		}
	st_case_45:
//line parser.go:940
		if data[(p)] == 101 {
			goto st46
		}
		goto st1
	st46:
		if (p)++; (p) == (pe) {
			goto _test_eof46
		}
	st_case_46:
		if data[(p)] == 98 {
			goto st47
		}
		goto st1
	st47:
		if (p)++; (p) == (pe) {
			goto _test_eof47
		}
	st_case_47:
		switch data[(p)] {
		case 32:
			goto tr20
		case 114:
			goto st48
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st48:
		if (p)++; (p) == (pe) {
			goto _test_eof48
		}
	st_case_48:
		if data[(p)] == 117 {
			goto st49
		}
		goto st1
	st49:
		if (p)++; (p) == (pe) {
			goto _test_eof49
		}
	st_case_49:
		if data[(p)] == 97 {
			goto st50
		}
		goto st1
	st50:
		if (p)++; (p) == (pe) {
			goto _test_eof50
		}
	st_case_50:
		if data[(p)] == 114 {
			goto st51
		}
		goto st1
	st51:
		if (p)++; (p) == (pe) {
			goto _test_eof51
		}
	st_case_51:
		if data[(p)] == 121 {
			goto st40
		}
		goto st1
	tr4:
//line parser.rl:20
		tok = p

		goto st52
	st52:
		if (p)++; (p) == (pe) {
			goto _test_eof52
		}
	st_case_52:
//line parser.go:1016
		switch data[(p)] {
		case 97:
			goto st53
		case 117:
			goto st55
		}
		goto st1
	st53:
		if (p)++; (p) == (pe) {
			goto _test_eof53
		}
	st_case_53:
		if data[(p)] == 110 {
			goto st54
		}
		goto st1
	st54:
		if (p)++; (p) == (pe) {
			goto _test_eof54
		}
	st_case_54:
		switch data[(p)] {
		case 32:
			goto tr20
		case 117:
			goto st49
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st55:
		if (p)++; (p) == (pe) {
			goto _test_eof55
		}
	st_case_55:
		switch data[(p)] {
		case 108:
			goto st56
		case 110:
			goto st57
		}
		goto st1
	st56:
		if (p)++; (p) == (pe) {
			goto _test_eof56
		}
	st_case_56:
		switch data[(p)] {
		case 32:
			goto tr20
		case 121:
			goto st40
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st57:
		if (p)++; (p) == (pe) {
			goto _test_eof57
		}
	st_case_57:
		switch data[(p)] {
		case 32:
			goto tr20
		case 101:
			goto st40
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	tr5:
//line parser.rl:20
		tok = p

		goto st58
	st58:
		if (p)++; (p) == (pe) {
			goto _test_eof58
		}
	st_case_58:
//line parser.go:1101
		if data[(p)] == 97 {
			goto st59
		}
		goto st1
	st59:
		if (p)++; (p) == (pe) {
			goto _test_eof59
		}
	st_case_59:
		switch data[(p)] {
		case 32:
			goto tr20
		case 114:
			goto st60
		case 121:
			goto st40
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st60:
		if (p)++; (p) == (pe) {
			goto _test_eof60
		}
	st_case_60:
		switch data[(p)] {
		case 32:
			goto tr20
		case 99:
			goto st61
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st61:
		if (p)++; (p) == (pe) {
			goto _test_eof61
		}
	st_case_61:
		if data[(p)] == 104 {
			goto st40
		}
		goto st1
	tr6:
//line parser.rl:20
		tok = p

		goto st62
	st62:
		if (p)++; (p) == (pe) {
			goto _test_eof62
		}
	st_case_62:
//line parser.go:1158
		if data[(p)] == 111 {
			goto st63
		}
		goto st1
	st63:
		if (p)++; (p) == (pe) {
			goto _test_eof63
		}
	st_case_63:
		if data[(p)] == 118 {
			goto st64
		}
		goto st1
	st64:
		if (p)++; (p) == (pe) {
			goto _test_eof64
		}
	st_case_64:
		switch data[(p)] {
		case 32:
			goto tr20
		case 101:
			goto st65
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st65:
		if (p)++; (p) == (pe) {
			goto _test_eof65
		}
	st_case_65:
		if data[(p)] == 109 {
			goto st66
		}
		goto st1
	st66:
		if (p)++; (p) == (pe) {
			goto _test_eof66
		}
	st_case_66:
		if data[(p)] == 98 {
			goto st67
		}
		goto st1
	st67:
		if (p)++; (p) == (pe) {
			goto _test_eof67
		}
	st_case_67:
		if data[(p)] == 101 {
			goto st68
		}
		goto st1
	st68:
		if (p)++; (p) == (pe) {
			goto _test_eof68
		}
	st_case_68:
		if data[(p)] == 114 {
			goto st40
		}
		goto st1
	tr7:
//line parser.rl:20
		tok = p

		goto st69
	st69:
		if (p)++; (p) == (pe) {
			goto _test_eof69
		}
	st_case_69:
//line parser.go:1234
		if data[(p)] == 99 {
			goto st70
		}
		goto st1
	st70:
		if (p)++; (p) == (pe) {
			goto _test_eof70
		}
	st_case_70:
		if data[(p)] == 116 {
			goto st71
		}
		goto st1
	st71:
		if (p)++; (p) == (pe) {
			goto _test_eof71
		}
	st_case_71:
		switch data[(p)] {
		case 32:
			goto tr20
		case 111:
			goto st66
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	tr8:
//line parser.rl:20
		tok = p

		goto st72
	st72:
		if (p)++; (p) == (pe) {
			goto _test_eof72
		}
	st_case_72:
//line parser.go:1274
		if data[(p)] == 101 {
			goto st73
		}
		goto st1
	st73:
		if (p)++; (p) == (pe) {
			goto _test_eof73
		}
	st_case_73:
		if data[(p)] == 112 {
			goto st74
		}
		goto st1
	st74:
		if (p)++; (p) == (pe) {
			goto _test_eof74
		}
	st_case_74:
		switch data[(p)] {
		case 32:
			goto tr20
		case 116:
			goto st75
		}
		if 9 <= data[(p)] && data[(p)] <= 13 {
			goto tr20
		}
		goto st1
	st75:
		if (p)++; (p) == (pe) {
			goto _test_eof75
		}
	st_case_75:
		if data[(p)] == 101 {
			goto st65
		}
		goto st1
	tr9:
//line parser.rl:20
		tok = p

		goto st76
	st76:
		if (p)++; (p) == (pe) {
			goto _test_eof76
		}
	st_case_76:
//line parser.go:1323
		if data[(p)] == 99 {
			goto st64
		}
		goto st1
	st_out:
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof

	_test_eof:
		{
		}
		if (p) == eof {
			switch cs {
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76:
//line parser.rl:28
				event.SetMessage(data[tok:p])

//line parser.go:1414
			}
		}

	}

//line parser.rl:72
}
