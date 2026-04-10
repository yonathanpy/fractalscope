# fractalscope

fractalscope is a multi-runtime telemetry processing framework designed for structured event ingestion, cross-language correlation, and deterministic anomaly evaluation across heterogeneous execution layers.

## System Architecture

The system is partitioned into isolated execution domains:

- go-stream: event ingestion, buffering, and runtime scheduling  
- rust-core: deterministic correlation engine and signal aggregation  
- python-analysis: statistical evaluation and anomaly scoring  
- c-syslayer: low-level system probe interface and execution hooks  
- bridge: serialization and cross-runtime data normalization  
- control: runtime supervision and fault containment  
- orchestrator: unified execution pipeline controller  

## Execution Model

The framework operates on a staged processing pipeline:

1. Event generation in Go runtime layer  
2. Serialization into structured transport format  
3. Python-based statistical evaluation and scoring  
4. Rust-based correlation and signal synthesis  
5. C-level validation of system-level signals  
6. Aggregation into unified output state  

Each stage operates independently and communicates via structured payloads.

## Build Process

### Rust Core
cargo build --release --manifest-path rust-core/Cargo.toml

### Go Stream Layer
go build ./go-stream

### Python Analysis Layer
python3 -m compileall python-analysis

### C System Layer
gcc c-syslayer/sys_probe.c -o sys_probe

## Runtime Execution

Primary execution entry:

go run orchestrator/main_pipeline.go

## Data Contract

All inter-layer communication follows a structured event schema:

- type: event classification tag  
- value: numeric or categorical signal  
- timestamp: ingestion time reference  
- metadata: optional contextual payload  

## Design Properties

- deterministic execution flow  
- bounded buffer processing model  
- stateless inter-stage communication  
- modular language-isolated runtime layers  
- explicit event transformation pipeline  

## Scope

fractalscope is intended for telemetry analysis, behavioral signal correlation, and multi-runtime evaluation in controlled system environments.
