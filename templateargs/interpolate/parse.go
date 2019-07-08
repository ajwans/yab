// Code generated by ragel
// @generated
package interpolate

import "fmt"

const interpolate_start int = 8
const interpolate_first_final int = 8
const interpolate_error int = 0

const interpolate_en_main int = 8


// Parse parses a string for interpolation.
//
// Variables may be specified anywhere in the string in the format ${foo} or
// ${foo:default} where 'default' will be used if the variable foo was unset.
func Parse(data string) (out String, _ error) {
	var (
		// Ragel variables
		cs  = 0
		p   = 0
		pe  = len(data)
		eof = pe

		// The following variables are used by us to build String up.

		// Index in data where the currently captured string started.
		idx int

		// Variable currently being built.
		v variable

		// Literal currently being read.
		l literal

		// Last read term (variable or literal) which we will append to the
		// output.
		t term
	)

	{
		cs = interpolate_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 1:
			goto st_case_1
		case 10:
			goto st_case_10
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 0:
			goto st_case_0
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		}
		goto st_out
	st_case_8:
		switch data[p] {
		case 36:
			goto st1
		case 92:
			goto st2
		}
		goto tr12
	tr12:
		idx = p
		l = literal(data[idx : p+1])
		t = l
		goto st9
	tr15:
		l = literal(data[idx : p+1])
		t = l
		goto st9
	tr18:
		out = append(out, t)
		idx = p
		l = literal(data[idx : p+1])
		t = l
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 36:
			goto tr16
		case 92:
			goto tr17
		}
		goto tr15
	tr16:
		out = append(out, t)
		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 123 {
			goto st3
		}
		goto tr0
	tr0:
		l = literal(data[p-1 : p+1])
		t = l
		goto st10
	tr2:
		l = literal(data[p : p+1])
		t = l
		goto st10
	tr8:
		t = v
		goto st10
	tr10:
		idx = p
		t = v
		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 36:
			goto tr16
		case 92:
			goto tr17
		}
		goto tr18
	tr17:
		out = append(out, t)
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		goto tr2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 95 {
			goto tr3
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr3
			}
		case data[p] >= 65:
			goto tr3
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr3:
		idx = p
		v.Name = data[idx : p+1]
		goto st4
	tr6:
		v.Name = data[idx : p+1]
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 58:
			goto tr7
		case 95:
			goto tr6
		case 125:
			goto tr8
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st5
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr6
				}
			case data[p] >= 65:
				goto tr6
			}
		default:
			goto tr6
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 95 {
			goto tr6
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr6
			}
		default:
			goto tr6
		}
		goto st0
	tr7:
		v.HasDefault = true
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 125 {
			goto tr10
		}
		goto tr9
	tr9:
		idx = p
		v.Default = data[idx : p+1]
		goto st7
	tr11:
		v.Default = data[idx : p+1]
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 125 {
			goto tr8
		}
		goto tr11
	st_out:
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof10:
		cs = 10
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

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 9, 10:
				out = append(out, t)
			}
		}

	_out:
		{
		}
	}


	if cs < 8 {
		return out, fmt.Errorf("cannot parse string %q", data)
	}

	return out, nil
}
