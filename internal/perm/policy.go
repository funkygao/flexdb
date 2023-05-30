package perm

type Policy struct {
	Action    []string    // [secretmanager.PutSecretValue, secretmanager.UpdateSecret]
	Resource  string      // *
	Condition interface{} //
	Effect    string      // allow
}

// Condition
/*
{
	"StringEquals": [
		{"secretmanager.ResourceTag/project": "${aws.PrincipleTag/project"},
	],
	"xxx": []
}
*/
