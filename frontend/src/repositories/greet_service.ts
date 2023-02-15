import { createGreetClient } from './grpc/client'

export const greeting = async (name: string): Promise<string> => {
  const client = createGreetClient()
  const response = await client.greet({
    name
  })
  return response.greeting
}
