#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#ifdef linux
#include <unistd.h>
#endif

#ifdef _WIN32
#include <windows.h>
#endif

#define BUFSIZE 128
#define infinite for(;;)  

char * fetch_output(char *command);
int get_active_container();
void dispatch_exceptions_frm_dckr(char *console_output);
void output_trace(char * stck, char * cnt_id, char * typ);

int main(int argc, char **argv)
{
    printf("--- Start Madlog ---\n");
	printf("Written by S. Kalski\n\n\n");
	get_active_container();
    return 0;
}

int get_active_container(){
	infinite
	{  
		char * returned_str = fetch_output("docker ps -q");
		char* pch = NULL;

		if(strlen(returned_str) == 0){
			printf("\n\nYour Docker is hidden or not available. Please make yourself a sudoer or reinstall docker.\n");
			printf("Maybe your Docker has not started yet or has no active containers running.\n");
			return 1;
		}

		pch = strtok(returned_str, "\r\n");
		while (pch != NULL){
			dispatch_exceptions_frm_dckr(pch);
			pch = strtok(NULL, "\n");
		}
		#ifdef _WIN32
			Sleep(2000);
		#endif
		#ifdef linux
			sleep(2);
		#endif

	} 

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
	pclose(fp);
    return output;
}

void dispatch_exceptions_frm_dckr(char *container_id){
	char * docker_cmd = "docker logs --since=2s ";
	char * command = malloc (strlen (docker_cmd) + strlen (container_id) + 1);;

	strcat(command, docker_cmd);
	strcat(command, container_id);
	char * returned_str = fetch_output(command);

	if(strstr(returned_str, "error") != NULL)
		output_trace(returned_str, container_id, "error");
	else if (strstr(returned_str, "stacktrace") != NULL)
		output_trace(returned_str, container_id, "error");
	
	return;
}

void output_trace(char * stck, char * cnt_id, char * typ){
		printf("Container %s returned a %s:\n", cnt_id, typ);
		printf("\n\n---- START OF LOG ----\n");
		printf("%s", stck);
		printf("\n\n---- END OF LOG ----\n");
}
