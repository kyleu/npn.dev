namespace system {
  class Cache {
    profile?: profile.Profile;

    public getProfile() {
      if (!this.profile) {
        throw "no active profile";
      }
      return this.profile;
    }

    apply(sj: MsgConnected) {
      system.cache.profile = sj.profile;
    }
  }

  export const cache = new Cache();
}
