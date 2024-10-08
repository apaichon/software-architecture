package main

import (
	"fmt"
	"plugin"
)

// Plugin interface
type Plugin interface {
	Execute() error
}

// TicketingKernel represents the core system
type TicketingKernel struct {
	plugins map[string]Plugin
}

func NewTicketingKernel() *TicketingKernel {
	return &TicketingKernel{
		plugins: make(map[string]Plugin),
	}
}

// LoadPlugin loads a plugin dynamically
func (k *TicketingKernel) LoadPlugin(name, path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("could not open plugin %s: %v", name, err)
	}

	symPlugin, err := p.Lookup("Plugin")
	if err != nil {
		return fmt.Errorf("could not find Plugin symbol in %s: %v", name, err)
	}

	plugin, ok := symPlugin.(Plugin)
	if !ok {
		return fmt.Errorf("unexpected type from symbol Plugin in %s", name)
	}

	k.plugins[name] = plugin
	return nil
}

// ExecutePlugin executes a loaded plugin
func (k *TicketingKernel) ExecutePlugin(name string) error {
	plugin, exists := k.plugins[name]
	if !exists {
		return fmt.Errorf("plugin %s not found", name)
	}
	return plugin.Execute()
}

// Example plugin (this would be in a separate file and compiled as a plugin)
type ExamplePlugin struct{}

func (p *ExamplePlugin) Execute() error {
	fmt.Println("Executing example plugin")
	return nil
}

// var Plugin ExamplePlugin

// Usage
func main() {
	kernel := NewTicketingKernel()

	// Load plugins
	err := kernel.LoadPlugin("example", "./example_plugin.so")
	if err != nil {
		fmt.Printf("Error loading plugin: %v\n", err)
		return
	}

	// Execute plugin
	err = kernel.ExecutePlugin("example")
	if err != nil {
		fmt.Printf("Error executing plugin: %v\n", err)
		return
	}
}
