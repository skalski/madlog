#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

#ifdef linux
#include <unistd.h>
#endif

#ifdef _WIN32
#include <windows.h>
#endif

#define BUFSIZE 256
#define infinite for(;;)  

char * fetch_output(char *command);
int get_active_container(char *arg);
void dispatch_exceptions_frm_dckr(char *console_output, char *arg);
void print_log_output(char *command, char *container_id);

int main(int argc, char *argv[])
{
    printf("--- Start MadLog ---\n");
	printf("Written by S. Kalski\n\n\n");
	if(argv[1] == NULL)
		get_active_container("");
	else
		get_active_container(argv[1]);
    return 0;
}

int get_active_container(char *arg){
	char * docker_cmd = " docker ps -q";
	char * command = malloc(strlen (arg) + strlen (docker_cmd)  + 1);

	strcat(command, arg);
	strcat(command, docker_cmd);

	char * returned_str = fetch_output(command);

	infinite
	{
		char* pch = NULL;
		if(strlen(returned_str) == 0){
			printf("\n\nYour Docker is hidden or not available. Please make yourself a sudoer or reinstall docker.\n");
			printf("Maybe your Docker has not started yet or has no active containers running.\n");
			printf("You can run `./madlog sudo` to ship around this issue.\n");
			return 1;
		}
		pch = strtok(returned_str, "\r\n");
		while (pch != NULL){
			dispatch_exceptions_frm_dckr(pch, arg);
			pch = strtok(NULL, "\n");
		}

		free(pch);

		#ifdef _WIN32
			Sleep(2000);
		#endif
		#ifdef linux
			sleep(2);
		#endif
	}

	printf("Oh dear, something went wrong or exit! %s\n", strerror(errno));
	return 0;
}

char * fetch_output(char *command){
    char buf[BUFSIZE];
	char * output;
	output = malloc(sizeof(char)*100);
    FILE *fp;

    if ((fp = popen(command, "r")) == NULL) {
        printf("Error opening pipe!\n");
        return "";
    }

    while (fgets(buf, BUFSIZE, fp) != NULL) {
		strcat(output, buf);
    }

	if (fp != NULL)
    {
		pclose(fp);
	}
    return output;
}

void print_log_output(char *command, char *container_id){
    FILE *fd;
    fd = popen(command, "r");
    if (!fd){
		printf("Error: %s %s", "can't fetch logs from ", container_id);
	}
 
    char  buffer[BUFSIZE];
    size_t chread;
    size_t comalloc = BUFSIZE;
    size_t comlen   = 0;
    char  *comout   = malloc(comalloc);
 
    while ((chread = fread(buffer, 1, sizeof(buffer), fd)) != 0) {
        if (comlen + chread >= comalloc) {
            comalloc *= 2;
            comout = realloc(comout, comalloc);
        }
        memmove(comout + comlen, buffer, chread);
        comlen += chread;
    }
 
    fwrite(comout, 1, comlen, stdout);
    free(comout);
    pclose(fd);
}

void dispatch_exceptions_frm_dckr(char *container_id, char *arg){
	char * docker_cmd = " docker logs --since=2s ";
	char * log_command[100];

	strcat(log_command, arg);
	strcat(log_command, docker_cmd);
	strcat(log_command, container_id);

	print_log_output(log_command, container_id);
	memset(log_command, 0, 100);
}
