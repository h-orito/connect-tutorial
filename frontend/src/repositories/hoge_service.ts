import { createHogeClient } from './grpc/client'

export const getHoge = async (id: string): Promise<Hoge | null> => {
  const client = createHogeClient()
  const response = await client.getHoge({
    id
  })
  if (!response.hoge) return null
  return {
    fuga: {
      piyo: response.hoge.fuga?.piyo ?? '???'
    } as Fuga
  } as Hoge
}

type Hoge = {
  fuga: Fuga
}

type Fuga = {
  piyo: string
}
