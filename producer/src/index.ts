import amqp from 'amqplib';
import cuid from 'cuid';
import faker from 'faker';
import initializeAMQP from './amqp';
import { QUEUE_NAMES } from './config';

interface Message {
  id: string;
  message: string;
}

const MESSAGES_TO_SEND = 100;
const MS_BETWEEN_MESSAGES = 200;

let connection: amqp.Connection;
let channel: amqp.Channel;

function sleep(ms: number): Promise<void> {
  return new Promise(res => {
    setTimeout(res, ms);
  });
}

function buildMessage(): Message {
  return { id: cuid(), message: faker.random.words() };
}

async function produce(): Promise<void> {
  console.log('Starting producer...');

  try {
    ({ connection, channel } = await initializeAMQP());
  } catch (err) {
    console.error(err);
    process.exit(1);
  }

  let messagesSent = 0;
  while (messagesSent < MESSAGES_TO_SEND) {
    const message = buildMessage();
    console.log(`Sending message with id ${message.id}...`);
    channel.sendToQueue(QUEUE_NAMES.demo, Buffer.from(JSON.stringify(message)));
    messagesSent += 1;
    await sleep(MS_BETWEEN_MESSAGES);
  }

  await channel.close();
  await connection.close();
}

produce();
