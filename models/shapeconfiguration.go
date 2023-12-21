package models

type ShapeConfiguration struct {
	Shape *Configuration
	Frame *Configuration
}

type ShapeConfigurationInspector struct {
	Shape *ConfigurationInspector
	Frame *ConfigurationInspector
}

func (sc *ShapeConfiguration) Inspect() (sci *ShapeConfigurationInspector) {
	sci = &ShapeConfigurationInspector{
		Shape: sc.Shape.Inspect(),
		Frame: sc.Frame.Inspect(),
	}
	return
}
