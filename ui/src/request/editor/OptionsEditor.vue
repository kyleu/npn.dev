<template>
  <div>
    <div class="mt">
      <label class="uk-form-label">
        Timeout (seconds)
        <input v-model.number="opts.timeout" class="uk-input" name="opt-timeout" type="number" />
      </label>
    </div>
    <div class="mt">
      <label class="uk-form-label">Ignore</label>
      <div>
        <label class="mr"><input v-model="opts.ignoreRedirects" type="checkbox" name="opt-ignoreRedirects" /> Redirects</label>
        <label class="mr"><input v-model="opts.ignoreReferrer" type="checkbox" name="opt-ignoreReferrer" /> Referrer</label>
        <label class="mr"><input v-model="opts.ignoreCerts" type="checkbox" name="opt-ignoreCerts" /> Certs</label>
        <label class="mr"><input v-model="opts.ignoreCookies" type="checkbox" name="opt-ignoreCookies" /> Cookies</label>
      </div>
    </div>
    <div class="mt">
      <label class="uk-form-label">
        Exclude Default Headers
        <input type="text" class="uk-input" name="opt-excludeDefaultHeaders" :value="opts.excludeDefaultHeaders" @input="opts.excludeDefaultHeaders = $event.target.value.split(',')" />
      </label>
    </div>
    <div class="mt">
      <label class="uk-form-label">
        Read Cookie Jars
        <input class="uk-input" name="opt-readCookieJars" type="text" :value="opts.readCookieJars" @input="opts.readCookieJars = $event.target.value.split(',')" />
      </label>
    </div>
    <div class="mt">
      <label class="uk-form-label">
        Write Cookie Jar
        <input v-model="opts.writeCookieJar" class="uk-input" name="opt-writeCookieJar" type="text" />
      </label>
    </div>
    <div class="mt">
      <label class="uk-form-label">
        SSL Cert
        <input v-model="opts.sslCert" class="uk-input" name="opt-sslCert" type="text" />
      </label>
    </div>
    <div class="mt">
      <label class="uk-form-label">
        User Agent Override
        <input v-model="opts.userAgentOverride" class="uk-input" name="opt-userAgentOverride" type="text" />
      </label>
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {Options} from "@/request/model";
import {requestEditingRef} from "@/request/state";

@Component
export default class OptionsEditor extends Vue {
  get opts(): Options {
    return requestEditingRef.value?.prototype?.options || {};
  }
}
</script>
