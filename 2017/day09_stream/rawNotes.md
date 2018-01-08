# Raw Notes

? What is the total score for all groups in your input?
{} is a group
<> is garbage
! deletes the char right after it

scoring
starting at 1, if stack is !empty, add previousNum + 1 to the count
{{},{}}
{ = 1, n=1 -- 1
		{{ = 1 + (1+1), if n=2 n-1 + n -- 1 + 2
	{{} becomes {, when deleted stack, move back to n-1 = n = 1 --- 1 + 2 -- don't add when closing
		{{ = 1 + (1+1), if n=2 n-1 + n -- 1 + 2 + 2
	{{} becomes {, when deleted stack, move back to n-1 = n = 1 -- don't add when closing
{} becomes = 0, n=0 -- dont add when closing
final == 1 + 2 + 2 == 5

{{{}}}, score of 1 + 2 + 3 = 6.

{ = 1
	{{ = 1 + 2
		{{{ = 1 + 2 + 3
		closing don't add
	closing don't add
closing don't add

{{<!!>},{<!!>},{<!!>},{<!!>}}, score of 1 + 2 + 2 + 2 + 2 = 9.
would become
{{<!>},{<!>},{<!>},{<!>}}
{ = 1
	{{ = 1 + 2
  ignore <
  ignore !
  ignore >
{{} closing
	{{ = 1 + 2 + 2
  ignore <
  ignore !
  ignore >
  etc

  {{<a!>},{<a!>},{<a!>},{<ab>}}, score of 1 + 2 = 3.
would become
{{<a!},{<a!},{<a!},{<ab>}}
{ = 1
{{ = 1 + 2
  < start tracking if found '>', nothing counts until then
a!},{<a!},{<a!},{<ab all don't count SKIP/CONTINUE
until '>' comes
{{} close
{{}} close, so 1 + 2 = 3



