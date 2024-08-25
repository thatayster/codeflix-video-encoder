#!/bin/bash

rabbitmqadmin declare exchange name=dlx type=fanout -u $RABBITMQ_DEFAULT_USER -p $RABBITMQ_DEFAULT_PASS
rabbitmqadmin declare queue name=videos-failed -u $RABBITMQ_DEFAULT_USER -p $RABBITMQ_DEFAULT_PASS
rabbitmqadmin declare binding source=dlx destination=videos-failed -u $RABBITMQ_DEFAULT_USER -p $RABBITMQ_DEFAULT_PASS
rabbitmqadmin declare queue name=videos-result -u $RABBITMQ_DEFAULT_USER -p $RABBITMQ_DEFAULT_PASS
rabbitmqadmin declare binding source=amq.direct destination=videos-result routing_key=jobs -u $RABBITMQ_DEFAULT_USER -p $RABBITMQ_DEFAULT_PASS