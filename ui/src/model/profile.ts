export default interface Profile {
  readonly userID: string;
  readonly name: string;
  readonly role: string;
  readonly theme: string;
  readonly navColor: string;
  readonly linkColor: string;
  readonly picture: string;
  readonly locale: string;
}
