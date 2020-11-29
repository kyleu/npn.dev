<template>
  <div>
    <select v-model="auth.type" class="uk-select">
      <option value="">no authentication</option>
      <option v-for="a in auths()" :key="a.key" :value="a.key">{{ a.title }}</option>;
    </select>

    <div v-if="auth.type === ''" class="mt">
      No auth!
    </div>
    <div v-else-if="auth.type === 'basic'" class="mt">
      <div>
        <label class="uk-form-label">
          Username
          <input v-model="auth.basicContent.username" class="uk-input" name="username" type="text" data-lpignore="true" />
        </label>
      </div>
      <div class="mt">
        <label class="uk-form-label">
          Password
          <input v-model="auth.basicContent.password" class="uk-input" name="password" type="text" data-lpignore="true" />
        </label>
      </div>
      <div class="mt">
        <label class="uk-form-label">Show Password</label>
        <div>
          <label>
            <input v-model="auth.basicContent.showPassword" name="showPassword" type="radio" value="true" class="uk-radio" />
            True
          </label>
          <label>
            <input v-model="auth.basicContent.showPassword" name="showPassword" type="radio" value="false" class="uk-radio" />
            False
          </label>
        </div>
      </div>
    </div>
    <div v-else class="mt">
      Unhandled [{{ auth.type }}] auth editor
    </div>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {AllTypes, AuthType} from "@/auth/model";
import {AuthConfig, authConfigRef} from "@/auth/state";

@Component
export default class AuthEditor extends Vue {
  get auth(): AuthConfig | undefined {
    return authConfigRef.value;
  }

  auths(): AuthType[] {
    return AllTypes.filter(t => !t.hidden);
  }
}
</script>
