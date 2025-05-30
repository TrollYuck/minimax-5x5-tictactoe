SOURCES = $(wildcard *.c)
BINARY = $(patsubst %.c, %, $(SOURCES))

FLAGS = -W -pedantic -std=c2x -O3
LIBS = -lgsl -lgslcblas -lm
.PHONY = all clean

all: $(BINARY)

$(BINARY): %: %.c
	$(CC) $(FLAGS) $^ -o $@ $(LIBS)
	strip $@

clean:
	rm -f $(BINARY)

