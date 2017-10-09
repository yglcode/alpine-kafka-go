Sample ping-pong app for bouncing msgs thru Kafka topic
=======================================================

1. runs with docker-compose (or swarm if you have swarm created)

2. git clone http://yglcode/alpine-kafka-go

3. cd alpine-kafka-go/sameples/pingpong

4. "docker-compose build":

        which produce image "yglcode/pingpong:kafka"
	"docker-compose push" if you use swarm and private registry.

5. add network for kafka:

	compose: docker network create -d bridge kafka-net
	swarm:   docker network create -d overlay kafka-net

6. bring up kafka-cluster(1 zookeeper, 1 kafka):

   	compose: docker-compose -f docker-compose.kafka.yml up -d
	swarm: docker stack deploy -c docker-compose.kafka.yml kafka-cluster

7. bring up pinger & ponger:

   	compose: docker-compose -f docker-compose.yml up -d
	swarm:   docker stack deploy -c docker-compose.yml pingpong

8. check messaging running:

   	compose: docker-compose -f docker-compose.yml logs -f
	swarm:   docker service logs -f pingpong_pinger_1

9. scale pinger or ponger:

         compose: docker-compose -f docker-compose.yml scale pinger=2
         swarm:   docker service scale pingpong_pinger=2

