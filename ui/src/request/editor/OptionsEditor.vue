<template>
  <div>
    <div class="mt">
      <label class="uk-form-label">
        Timeout
        <input v-model="opts.timeout" class="uk-input" name="opt-timeout" type="number" />
      </label>
    </div>
    <div class="mt">
      <label class="uk-form-label">Ignore</label>
      <div>
        <label class="uk-margin-right"><input v-model="opts.ignoreRedirects" type="checkbox" name="opt-ignoreRedirects" /> Redirects</label>
        <label class="uk-margin-right"><input v-model="opts.ignoreReferrer" type="checkbox" name="opt-ignoreReferrer" /> Referrer</label>
        <label class="uk-margin-right"><input v-model="opts.ignoreCerts" type="checkbox" name="opt-ignoreCerts" /> Certs</label>
        <label class="uk-margin-right"><input v-model="opts.ignoreCookies" type="checkbox" name="opt-ignoreCookies" /> Cookies</label>
      </div>
    </div>
    <div class="mt">
      <label class="uk-form-label">
        Exclude Default Headers
        <input v-model="opts.excludeDefaultHeaders" class="uk-input" name="opt-excludeDefaultHeaders" type="text" />
      </label>
    </div>
    <div class="mt">
      <label class="uk-form-label">
        Read Cookie Jars
        <input v-model="opts.readCookieJars" class="uk-input" name="opt-readCookieJars" type="text" />
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
import {Component, Prop, Vue} from "vue-property-decorator";
import {Options} from "@/request/model";

@Component
export default class OptionsEditor extends Vue {
  @Prop() opts!: Options | undefined;
}
</script>
