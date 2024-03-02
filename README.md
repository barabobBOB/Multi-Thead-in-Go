# 멀티 스레드 그랩

용감한 쿠키는 평소 사용하는 파일 내 문자열 검색 툴 grep보다 더 빠른 속도를 자랑하는 ack나 ag같은 다양한 툴들이 있다는 것을 들었다.
자신만의 dgrep을 만들어 널리 세상을 이롭게 하고 싶다고 생각하게 된 용감한 쿠키는 그 첫 단계로 멀티스레드로 동작하는 grep을 직접 만들어보기로 결심하였다.
원하는 스펙에 맞는 파일 내 문자열 검색 툴을 만들어보자.

입력 형식
- dgrep {keyword} {relative path}

출력 형식
- 파일의 각 line에 keyword가 있는 경우, 해당 파일과 줄 번호를 출력한다.

조건
1. relative path가 디렉토리인 경우 디렉토리 내 모든 파일에 대해 검사를 진행한다.
2. relative path 내에 또 다른 디렉토리가 존재하는 경우, 각 디렉토리 내 모든 파일에 대한 검사 또한 진행한다.
3. 멀티 스레드를 이용하여 최대한 빠르게 작업을 완료하도록 작성한다.
4. 동일한 파일에 대한 검사 결과는 한 번에 출력되어야 한다.
5. Directory 내 symlink는 없다고 가정한다.
6. 파일들은 모두 UTF8 인코딩으로 작성된 Text파일이라고 가정한다.

출처: https://tech.devsisters.com/posts/server-position-coding-test/

구현 과정: https://velog.io/@chltpdus48/Multi-Thead-in-Go
