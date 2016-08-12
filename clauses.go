package cctools

//Creation of Clause struct and related methods; with JSON keys incorporated

type Clause struct {
	Name		string		`json:"name"`
	Heading		string		`json:"heading"`
	Text		string		`json:"text"`
	Description	string 		`json:"description"`
	Tags		[]string	`json:"tags"`
	Type 		string		`json:"type"`
	//Benefit		bool		`json:"benefit"` //This mainly applies to conditions; and if true, means one party is benefitted by condition
	//BftPty		Party		`json:"bftpty"` //If one or more parties benefit from a condition (to the exclusion of others), then BftPty identifies those parties
}

func ImportClause(filename string) Clause {
	data :=LoadFromJSON(filename)
	newClause :=JSONtoClause(data)
    return  newClause
}

//Setter for Name attribute

func (c *Clause) SetName(n string) {
	c.Name = n
}

//Setter for Heading attribute

func (c *Clause) SetHeading(h string) {
	c.Heading = h
}

//Setter for Text attribute

func (c *Clause) SetText(t string) {
	c.Text = t
}

//Setter for Description attribute

func (c *Clause) SetDescription(d string) {
	c.Description = d
}

func (c *Clause) AddTag(t string) {
	newTag := t
	oldTags := c.Tags
	newTags := append(oldTags, newTag) 
	c.Tags = newTags
}

//Setter for Type attribute

func (c *Clause) SetType(t string) {
	c.Type = t
}

/* Setter for Benefit attribute

func (c *Clause) SetBenefit(b bool) {
	c.Benefit = b
}

//Setter for Benefit attribute
//Generally, the benefitted party has the right to waive a condition and proceed with the transaction

func (c *Clause) SetBftPty(b Party) {
	c.BftPty = b
}
*/
//NOTE: This needs to be worked out because a changed clause will not be reflected in an existing documents because Clauses are a 
// type value rather than pointers to Clauses.  So the edited clause is not sharing data with the Clause that is an attribute
// in the Document object. 

func (c *Clause) EditText(t string) {
	c.Text = t
}


