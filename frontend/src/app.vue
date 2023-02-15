<template>
  <div>
    <input type="text" v-model="name" />
    <button @click="hoge">go</button>
    {{ res }}
  </div>
</template>

<script setup lang="ts">
import { greeting } from '~/repositories/greet_service'
import { getHoge } from '~/repositories/hoge_service'

const name = ref('Jane')
const res = ref('')

const greet = async () => {
  res.value = await greeting(name.value)
}
const hoge = async () => {
  const response = await getHoge(name.value)
  res.value = response?.fuga.piyo ?? 'not found'
}
</script>
