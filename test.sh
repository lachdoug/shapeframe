export PATH=$PATH:.

redColor='\033[0;31m'
greenColor='\033[0;32m'
yellowColor='\033[0;33m'
blueColor='\033[0;34m'
noColor='\033[0m'

errorCount=0
successCount=0
arg1="$1"

sf() {
    echo -e "${blueColor}sf $@${noColor}"
    command sf "$@"
    if [ $? != 0 ]
    then
        handleFailure
    else
        handleSuccess
    fi
}

handleFailure() {
    echo -e "${redColor}Fail${noColor}\n"
    let errorCount=errorCount+1
    if [ "$arg1" != "-continue" ]
    then
        exit 1
    fi    
}

handleSuccess() {
    echo -e "${greenColor}Pass${noColor}\n"
    let successCount=successCount+1
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

t() {
    echo -e "\n${yellowColor}$@${noColor}\n"   
}

t Nuking
sf nuke -confirm

t Adding and removing workspaces
sf a w W1
sf rm w W1
sf a w W1
sf en W1
sf rm w

t Adding and removing directories
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf rm d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf en W1
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf rm d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf rm w

t Adding and removing repositories
sf a w W1
sf a r -workspace W1 github.com/lachdoug/shapeframe-apps
sf rm r -workspace W1 github.com/lachdoug/shapeframe-apps
sf en W1
sf a r github.com/lachdoug/shapeframe-apps
sf rm r github.com/lachdoug/shapeframe-apps
sf a r -ssh github.com/lachdoug/shapeframe-apps
sf pull github.com/lachdoug/shapeframe-apps
sf rm w

t Adding and removing frames
sf a w W1
sf en W1
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf rm f docker-local
sf a f docker-local
sf en docker-local
sf rm f
sf ex
sf a f -workspace W1 docker-local
sf rm f -workspace W1 docker-local
sf rm w W1

t Adding and removing shapes
sf a w W1
sf en W1
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf en docker-local
sf a s apache
sf rm s apache
sf ex
sf a s -frame docker-local apache
sf rm s -frame docker-local apache
sf ex
sf a s -workspace W1 -frame docker-local apache
sf rm s -workspace W1 -frame docker-local apache
sf rm w W1

t Listing
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf ls w
sf ls f -all
sf ls s -all
sf ls sr -all
sf ls fr -all
sf ls f -workspace W1 
sf ls s -workspace W1 
sf ls sr -workspace W1 
sf ls fr -workspace W1 
sf en W1
sf ls f
sf ls s
sf ls sr
sf ls fr
sf rm w

t Inspecting
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf en W1
sf i w
sf rm w

t Contexting
sf x
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf en W1
sf x
sf en docker-local
sf x
sf en apache
sf x
sf rm w

t Configuring
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf c f -workspace W1 -frame docker-local blue foo
sf c s -workspace W1 -frame docker-local -shape apache blue foo cool
sf en W1
sf c f -frame docker-local blue foo
sf c s -frame docker-local -shape apache blue foo cool
sf en docker-local
sf c s -shape apache blue foo cool
sf rm w

t Labelling
sf a w -about "Workspace One" W1 
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf ll w -name W1! -about "Workspace One!" W1
sf ll f -workspace W1! -name docker-local! -about "Deploy containers on local Docker Engine!" docker-local

sf ll s -workspace W1! -frame docker-local! -name apache! -about "Apache (Web Server)!" apache
sf en W1!
sf ll w -name W1!! -about "Workspace One!!"
sf ll f -name docker-local!! -about "Deploy containers on local Docker Engine!!" docker-local!
sf ll s -frame docker-local!! -name apache!! -about "Apache (Web Server)!!" apache!
sf en docker-local!!
sf ll f -name docker-local!!! -about "Deploy containers on local Docker Engine!!!"
sf ll s -name apache!!! -about "Apache (Web Server)!!!" apache!!
sf en apache!!!
sf ll s -name apache!!!! -about "Apache (Web Server)!!!!"
sf rm w

t Orchestrating
sf a w -about "Workspace One" W1 
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf o -workspace W1 -frame docker-local
sf en W1
sf o -frame docker-local
sf en docker-local
sf o

t Helping
sf -help
sf list -help
sf ls workspaces -help
sf ls frames -help
sf ls shapes -help
sf ls framers -help
sf ls shapers -help
sf label -help
sf ll shape -help
sf ll frame -help
sf ll workspace -help
sf configure -help
sf c shape -help
sf c frame -help
sf inspect -help
sf add -help
sf a shape -help
sf a frame -help
sf a workspace -help
sf a repository -help
sf a directory -help
sf remove -help
sf rm shape -help
sf rm frame -help
sf rm workspace -help
sf rm repository -help
sf rm directory -help
sf pull -help
sf enter -help
sf exit -help
sf context -help
sf orchestrate -help
sf nuke -help

report
