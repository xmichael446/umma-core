<template>
  <div>
    <SpTheme>
      <SpNavbar
        :links="navbarLinks"
        :active-route="router.currentRoute.value.path"
      />
      <router-view />
    </SpTheme>
  </div>
</template>

<script lang="ts">
import { SpNavbar, SpTheme } from '@starport/vue'
import { computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

export default {
  components: { SpTheme, SpNavbar },

  setup() {
    // store
    let $s = useStore()

    // router
    let router = useRouter()

    // state
    let navbarLinks = [
      { name: 'Portfolio', url: '/portfolio' },
      { name: 'Data', url: '/data' }
    ]

    // computed
    let address = computed(() => $s.getters['common/wallet/address'])

    // lh
    onBeforeMount(async () => {
      // await $s.dispatch('common/env/init')
      await $s.dispatch('common/env/init',{
        apiNode: 'http://localhost:1317',
        rpcNode: 'http://localhost:26657',
        wsNode: 'ws://localhost:26657',
        addrPrefix: 'umma',
        sdkVersion: 'Stargate',
        getTXApi: 'http://localhost:26657/tx?hash=0x',
        offline: false,
        refresh: 5000
      })

      router.push('portfolio')
    })

    return {
      navbarLinks,
      // router
      router,
      // computed
      address
    }
  }
}
</script>

<style scoped lang="scss">
body {
  margin: 0;
}
</style>
