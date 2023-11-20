#!/bin/sh
while true; do
  case "$1" in
    -ud|--up-dev)
        sleep 2
        printf '\033[0;34m ******** setup on development ********\n \033[0m'
        docker-compose --env-file ../deploy/.env.dev -f ../deploy/docker-compose.yml up --build --detach

        sleep 1
        printf '\033[0;34m ******** migration setup database postgresql ********\n \033[0m'
        cd ../internal/repository/postgresql/migrations || exit
        sql-migrate up

           exit 0
            ;;
    -dd|--down-dev)
        sleep 2
        printf '\033[0;34m ******** docker down development ********\n \033[0m'
        docker-compose --env-file ../deploy/.env.dev -f ../deploy/docker-compose.yml down
      exit 0
      ;;
     -dvd|--down-volume-dev)
            sleep 2
            printf '\033[0;34m ******** docker down ********\n \033[0m'
            docker-compose --env-file ../deploy/.env.dev -f ../deploy/docker-compose.yml down
            sleep 3
            printf '\033[0;34m ******** remove volume dev ********\n \033[0m'
            docker volume rm deploy_vm_infiblog
          exit 0
          ;;
        -up|--up-prod)
            sleep 2
            printf '\033[0;34m ******** deploy prod ********\n \033[0m'
            docker-compose --env-file ../deploy/.env.prod -f ../deploy/docker-compose.yml up --build --detach
               exit 0
                ;;
        -dp|--down-prod)
            sleep 2
            printf '\033[0;34m ******** down prod ********\n \033[0m'
            docker-compose --env-file ../deploy/.env.prod -f ../deploy/docker-compose.yml down

               exit 0
                ;;
    --)
      break;;
     *)
      printf "sub command usage:  [-ud | --up-dev] [-up | --up-prod] [-dvd | --down-volume-dev] [-dd | --down-dev] [-dp | --down-prod] \n"
      exit 1;;
  esac
done
