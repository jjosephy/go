package dp

import (
    "errors"
    "fmt"
)

type SubSetSum struct {
    Set []int
    Sum int
    solution [][]bool
    List []int
}

func NewSubSetSum(a []int, s int) (*SubSetSum, error) {
    if len(a) == 0 {
        return nil, errors.New("invalid arg")
    }

    return &SubSetSum {
        Set : a,
        Sum : s,
    }, nil
}

func (s *SubSetSum) initializeTable(l int, x int) {
    s.solution = make([][]bool, l, l)

    for i := 0; i < l; i++ {
        s.solution[i] = make([]bool, x, x)
    }

    for i := 1; i < x; i++ {
        s.solution[0][i] = false
    }

    for i := 0; i < l; i++ {
        s.solution[i][0] = true
    }
}

func (s *SubSetSum) IsSubSetSum() (bool) {
    l := len(s.Set) + 1
    x := s.Sum + 1
    s.initializeTable(l, x)

    // build the table
    for i := 1; i <= len(s.Set); i++ {
      for j := 1; j <= s.Sum; j++ {
        s.solution[i][j] = s.solution[i-1][j];
        if ((s.solution[i][j] == false) && j >= s.Set[i-1]) {
            s.solution[i][j] = s.solution[i][j] || s.solution[i-1][j - s.Set[i-1]];
        }
      }
    }

    // find the path
    sum := 0
    z := len(s.Set)
    y := s.Sum

    for {
        if z == 0 || y == 0 {
            break;
        }

        if s.solution[z][y] == true && s.solution[z-1][y] == false {
            t := sum + s.Set[z-1]
            if t <= s.Sum {
                sum = t
                s.List = append(s.List, s.Set[z-1])
            }
            y--
        } else if s.solution[z][y] == true && s.solution[z-1][y] == true {
            z--
        } else {
            y--
        }
    }

    return s.solution[len(s.Set)][s.Sum]
}

func (s *SubSetSum) Dump() {

    x := s.Sum + 1
    l := len(s.Set)

    fmt.Println("")
    fmt.Println("Sum: ", s.Sum)
    for i := 0; i < x; i++ {
        fmt.Print(i, "   |")
    }
    fmt.Println("")
    for i := 0; i < l; i++ {
        if i == 0 {
            fmt.Print("0 - ")
        } else {
                fmt.Print(s.Set[i -1], " - ")
        }

        for j := 0; j < x; j++ {
            fmt.Print("|", s.solution[i][j])
        }
        fmt.Println("")
    }
    fmt.Println("")
    fmt.Println("Show list")
    for i := 0; i < len(s.List); i++ {
        fmt.Print(s.List[i], "  - ")
    }
    fmt.Println("")
}
