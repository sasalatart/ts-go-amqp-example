export const AMQP_URI =
  process.env.AMQP_URI || 'amqp://guest:guest@localhost:5672';

export const QUEUE_NAMES = {
  demo: 'ts-go-amqp-example',
};
