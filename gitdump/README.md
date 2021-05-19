# gitdump

**git** **dump**er

It prints all the matches in a git repository with a specified pattern.

This is a @tomnomnom tool (https://twitter.com/tomnomnom/status/1133345832688857095).

### Install

- `git clone https://github.com/edoardottt/lit-bb-hack-tools`
- `cd lit-bb-hack-tools/gitdump`
- `chmod +x gitdump`
- `sudo cp gitdump /usr/bin`

### Usage

- `(github-repo)>$ gitdump PATTERN-HERE`

### Sample output

```
(~/github/lit-bb-hack-tools)>$ gitdump func

func AddHeaders(input string) string {
func AddProtocol(input string) string {
func extractExtensions(input []string) {
func GetHeaders(target string, result map[string][]string) map[string][]string {
func GetOnlySubs(input string) string {
func GetProtocol(input string) string {
func main() {
func removeDuplicateValues(intSlice []string) []string {
func removeDuplicateValues(strSlice []string) []string {
func RemoveHeaders(input string) string {
func RemovePort(input string) string {
func RetrieveHeaders(input []string) {
func ScanFlag() bool {
func ScanTargets() []string {
(~/github/lit-bb-hack-tools)>$ gitdump func
suffice to ensure that the continued functioning of the modified object
```
