#! bin/bash

echo "////////////////////////////////////////////\n"
echo "./checker\n"
./checker &&
wait$1
echo "////////////////////////////////////////////\n"
echo '.</checker "0 one 2 3"\n'
./checker "0 one 2 3" &&
wait$1
echo "////////////////////////////////////////////\n"
echo 'echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"n'
echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5" &&
wait$1
echo "////////////////////////////////////////////\n"
echo 'echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"'
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2" &&
wait$1
echo "////////////////////////////////////////////\n"
echo 'ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"'
ARG="4 67 3 87 23"; ../push-swap/push-swap "$ARG" | ./checker "$ARG"

