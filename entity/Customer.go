package entity

import (
	"gopkg.in/gorp.v1"
	"time"
	"errors"
	"strings"
	. "github.com/newtechfellas/CurryPoint/util"
	"github.com/newtechfellas/CurryPoint"
)

type Customer struct {
	Email        string
	Password     string
	FirstName    string
	LastName     string
	PhoneNumber  string
	Address      struct {
			     Line1 string
			     Line2 string
			     City  string
			     State string
			     Zip   int
		     }
	CreatedTime  time.Time
	VerifiedTime time.Time
}

// implement the PreInsert and PreUpdate hooks
func (c *Customer) PreInsert(s gorp.SqlExecutor) error {
	c.CreatedTime = time.Now()
	var mandatoryColsMissing []string
	if c.Email == "" {
		mandatoryColsMissing = append(mandatoryColsMissing, "Email")
	}
	if c.Password == "" {
		mandatoryColsMissing = append(mandatoryColsMissing, "Password")
	}
	if len(mandatoryColsMissing) > 0 {
		return errors.New("Mandatory properties missing " + strings.Join(mandatoryColsMissing, ","))
	}
	return nil
}

func (c *Customer) Exists() bool {
	sql := "select count(*) from Customer where Email =?"
	count, err := CurryPoint.DBMAP.SelectInt(sql, c.Email)
	if err != nil {
		LogError("Failed to user existence using " + c.PhoneNumber + ". Error is " + err.Error())
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}

func (c *Customer) ConfigureDBMap(dbMap *gorp.DbMap) {
	tbl := dbMap.AddTableWithName(*c, "Customer").SetKeys(false, "Email")
	mandatoryCols := []string{"Password"}
	for _, col := range mandatoryCols() {
		tbl.ColMap(col).SetNotNull(true)
	}
}