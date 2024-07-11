package types

type EntityType struct {
	Name       string
	Attributes []EntityTypeAttributes
}

type EntityTypeAttributes struct {
	ObjType string `yaml:"type"`
	Config  EntityConfigType
}

type EntityConfigType struct {
	Description string
	Fields      map[string]EntityField `yaml:"fields,omitempty"`
}

type EntityField struct {
	FieldType        string `yaml:"type"`
	FieldDescription string `yaml:"descriptions"`
}

// Character:
//     type: object
//     heirs: # the opposite of « inherits »
//            # optional if « inherits » already exists on daughters classes
//       - CharacterWarrior
//       - CharacterWizard
//     config:
//         fields:
//             id: {type: Int!}
//             type: {type: String!}
//             name: {type: String!}
//             staminaPoints: {type: Int!}
