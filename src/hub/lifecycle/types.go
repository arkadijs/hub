package lifecycle

type Request struct {
	Verb                     string
	DryRun                   bool
	ManifestFilenames        []string
	StateFilenames           []string
	LoadFinalState           bool
	Component                string   // invoke
	Components               []string // deploy & undeploy, backup
	OffsetComponent          string   // deploy & undeploy
	LimitComponent           string   // deploy & undeploy
	GuessComponent           bool     // undeploy
	StrictParameters         bool
	OsEnvironmentMode        string
	EnvironmentOverrides     string
	ComponentsBaseDir        string
	PipeOutputInRealtime     bool
	GitOutputs               bool
	GitOutputsStatus         bool
	Environment              string
	StackInstance            string
	Application              string
	SaveStackInstanceOutputs bool
}
