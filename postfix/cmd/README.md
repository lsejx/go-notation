# Installation
	go install github.com/lsejx/go-notation/postfix/cmd/pofc

# Usage

### help
	pofc -h

### 64bit float mode
	pofc 1 1 +         // 1+1=2
	pofc pi 2 x        // 2π=6.28...
	pofc 2 8 xx        // pow(2,8)=256
	pofc 2 1 2 / xx    // sqrt(2)=1.41...
	pofc 1.234 1k x    // 1234
	pofc 1k 1m x       // 1

### 64bit signed int mode
	pofc -i 1 1 +
	pofc -i max
	pofc -i min

### 64bit unsigned int mode
	pofc -u 1 1 +
	pofc -u max
