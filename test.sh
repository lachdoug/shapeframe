export PATH=$PATH:.

redColor='\033[0;31m'
greenColor='\033[0;32m'
blueColor='\033[0;34m'
noColor='\033[0m'

errorCount=0
successCount=0

sf() {
    echo -e "${blueColor}sf $@${noColor}"
    command sf "$@"
    exitCode=$?
    if [ $exitCode -ne 0 ]
    then
        echo -e "${redColor}Fail${noColor}\n"
        let errorCount=errorCount+1
    else
        echo -e "${greenColor}Pass${noColor}\n"
        let successCount=successCount+1
    fi
}

report() {
    echo -e "${greenColor}Pass: ${successCount} ${redColor}Fail: ${errorCount}${noColor}"
    if [[ "$errorCount" -eq 0 ]]
    then
        echo -e "${greenColor}Tests passed${noColor}"
        exit 0
    else
        echo -e "${redColor}Tests failed${noColor}"
        exit 1
    fi
}

sf nuke <<END
Y
END
sf \?
sf cr w W1
sf l w
sf en W1
sf \?
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf i
sf r d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf i
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf i
sf l sr
sf l fr
sf cr f docker-local
sf l f
sf en docker-local
sf \?
sf cr s apache
sf l s
sf en apache
sf i
sf \?
sf a r github.com/lachdoug/shapeframe-apps
sf i
sf r r github.com/lachdoug/shapeframe-apps
sf a r --ssh github.com/lachdoug/shapeframe-apps
sf i
sf \?

report
