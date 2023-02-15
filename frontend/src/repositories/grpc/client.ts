import { ServiceType } from '@bufbuild/protobuf'
import {
  createPromiseClient,
  createConnectTransport,
  PromiseClient
} from '@bufbuild/connect-web'
import { GreetService } from './greet/v1/greet_connectweb'
import { HogeService } from './hoge/v1/hoge_connectweb'

const transport = createConnectTransport({
  baseUrl: 'http://localhost:8080'
})

export const createGreetClient = () => {
  return createClient(GreetService)
}

export const createHogeClient = () => {
  return createClient(HogeService)
}

const createClient = <T extends ServiceType>(service: T): PromiseClient<T> => {
  return createPromiseClient(service, transport)
}
