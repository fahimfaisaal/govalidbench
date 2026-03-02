bench:
	go test -bench=. -benchmem -run=none

bench-zog:
	@echo "======"
	@echo "Benchmarking zog..."
	@echo "======"
	go test ./benchmarks -bench=/zog/ -benchmem -run=none

bench-validator:
	@echo "======"
	@echo "Benchmarking validator..."
	@echo "======"
	go test ./benchmarks -bench=/validator/ -benchmem -run=none

bench-ozzo:
	@echo "======"
	@echo "Benchmarking ozzo..."
	@echo "======"
	go test ./benchmarks -bench=/ozzo/ -benchmem -run=none

bench-govalidator:
	@echo "======"
	@echo "Benchmarking govalidator..."
	@echo "======"
	go test ./benchmarks -bench=/govalidator/ -benchmem -run=none

bench-all:
	@echo "======"
	@echo "Running all benchmarks..."
	@echo "======"
	go test ./benchmarks -bench=. -benchmem -run=none
	@echo "======"
	@echo "All benchmarks completed"
	@echo "======"


# install vizh: go install github.com/goptics/vizb@latest
vizb-all:
	go test ./benchmarks -bench=. -benchmem -run=none | vizb -o charts/index.html -p n/x/y -s asc -n "Validation Library Benchmarks Comparison"