package commons

import "errors"

var (
	ErrLoanNotFound           = errors.New("loan not found")
	ErrIncompleteRepayment    = errors.New("incomplete pending repayment")
	ErrInvalidRepaymentAmount = errors.New("repayment amount must be exact for the week")
)
