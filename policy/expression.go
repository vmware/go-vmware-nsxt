package policy

// All the nodes of the expression extend from this abstract class. This is present for extensibility.
type Expression struct {
	// Common fields for all expression types
	ResourceType    string `json:"resource_type"`
	Id              string `json:"id,omitempty"`
	Path            string `json:"path,omitempty"`
	RelativePath    string `json:"relative_path,omitempty"`
	ParentPath      string `json:"parent_path,omitempty"`
	RemotePath      string `json:"remote_path,omitempty"`
	MarkedForDelete bool   `json:"marked_for_delete,omitempty"`
	Overridden      bool   `json:"overridden,omitempty"`
	Protection      string `json:"_protection,omitempty"`

	// Condition-specific fields
	MemberType    string `json:"member_type,omitempty"`
	Key           string `json:"key,omitempty"`
	Operator      string `json:"operator,omitempty"`
	ScopeOperator string `json:"scope_operator,omitempty"`
	Value         string `json:"value,omitempty"`

	// ConjunctionOperator-specific fields
	ConjunctionOperator string `json:"conjunction_operator,omitempty"`

	// PathExpression-specific fields
	Paths []string `json:"paths,omitempty"`
}

// Condition represents a condition expression
type Condition struct {
	Expression
	MemberType    string `json:"member_type,omitempty"`
	Key           string `json:"key,omitempty"`
	Operator      string `json:"operator,omitempty"`
	ScopeOperator string `json:"scope_operator,omitempty"`
	Value         string `json:"value,omitempty"`
}

// ConjunctionOperator represents a conjunction operator expression
type ConjunctionOperator struct {
	Expression
	ConjunctionOperator string `json:"conjunction_operator,omitempty"`
}

// PathExpression represents a path expression
type PathExpression struct {
	Expression
	Paths []string `json:"paths,omitempty"`
}

// Helper functions to create specific expression types
func UpdateCondition(group *Group, memberType, key, operator, scopeOperator, value string) {
	for ix := range group.Expression {
		if group.Expression[ix].ResourceType == "Condition" {
			// Update existing condition
			group.Expression[ix].MemberType = memberType
			group.Expression[ix].Key = key
			group.Expression[ix].Operator = operator
			group.Expression[ix].ScopeOperator = scopeOperator
			group.Expression[ix].Value = value
			return
		}
	}
	group.Expression = append(group.Expression, Expression{
		ResourceType:  "Condition",
		MemberType:    memberType,
		Key:           key,
		Operator:      operator,
		ScopeOperator: scopeOperator,
		Value:         value,
	})

}

func UpdateConjunctionOperator(group *Group, operator string) {
	for ix := range group.Expression {
		if group.Expression[ix].ResourceType == "ConjunctionOperator" {
			// Update existing conjunction operator
			group.Expression[ix].ConjunctionOperator = operator
			return
		}
	}
	group.Expression = append(group.Expression, Expression{
		ResourceType:        "ConjunctionOperator",
		ConjunctionOperator: operator,
	})

}

func UpdatePathExpression(group *Group, paths []string) {

	for ix := range group.Expression {
		if group.Expression[ix].ResourceType == "PathExpression" {
			// Update existing path expression
			group.Expression[ix].Paths = paths
			return
		}
	}
	group.Expression = append(group.Expression, Expression{
		ResourceType: "PathExpression",
		Paths:        paths,
	})
}
