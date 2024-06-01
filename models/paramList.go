package models

type ParamList []string

func (l *ParamList) String() string {
    return "Param list representation"
}

func (l *ParamList) Set(value string) error {
    *l = append(*l, value)
    return nil
}