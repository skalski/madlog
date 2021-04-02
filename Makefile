OBJS	= madlog.o
SOURCE	= madlog.c
HEADER	=
OUT	= madlog
CC	 = gcc
FLAGS	 = -g -c -Wall
LFLAGS	 =

all: $(OBJS)
	$(CC) -g $(OBJS) -o $(OUT) $(LFLAGS)

madlog.o: madlog.c
	$(CC) $(FLAGS) madlog.c -std=c99 -lcunit


clean:
	rm -f $(OBJS) $(OUT)

run: $(OUT)
	./$(OUT)
