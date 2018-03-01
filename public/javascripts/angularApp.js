var app = angular.module('mcfly', ['angularFileUpload']);

app.controller('MainCtrl', ['$scope','$timeout', 'FileUploader', function($scope, $timeout, FileUploader) {
	$scope.hello = "Angular";  

    // Check out the index.js router to see how the 
    // file upload is handled on the back end 
	uploader = $scope.uploader = new FileUploader({
		url: '/mcfly/upload'
	}); 

	uploader.filters.push({
        name: 'fileWhitelist',
        fn: function(item, options) {
            var fileCheck = false;
            switch (item.name.split('.').pop()) {
                case 'vcf':
                    fileCheck = true;
                    break;
                case 'zip':
                    fileCheck = true;
                    break;
                case 'doc':
                    fileCheck = true;
                    break;
                case 'docx':
                    fileCheck = true;
                    break;
                case 'xls':
                    fileCheck = true;
                    break;
                case 'xlsx':
                    fileCheck = true;
                    break;
                case 'gz':
                    fileCheck = true;
                    break;
                case 'txt':
                    fileCheck = true;
                    break;
                default:
                    break;

            }
            return fileCheck;
        }
    }, { 
        name: 'queueLength', 
        fn: function(item, options) {
            return this.queue.length < 8; 
        }
    });

    // Callbacks - Useful for debugging. Check out the documentation
    // (https://github.com/nervgh/angular-file-upload) for more. 

    uploader.autoUpload = true;
    uploader.onWhenAddingFileFailed = function(item, filter, options) {
        console.info('onWhenAddingFileFailed', item, filter, options);
        console.log (item.name.split('.').pop());
        $scope.doFade = false; 

        $scope.showError = true; 

        if (filter.name === "queueLength") {
            console.log("Too many items"); 
            $scope.errorMessage = "Queue capped at eight items."
        } else {
            console.log("Wrong file type");
            $scope.errorMessage = "Sorry, that file type isn't accepted.";
        }

        $timeout(function() {
            $scope.doFade = true; 
        }, 2500); 

    };

    uploader.onSuccessItem = function(fileItem, response, status, headers) {
        //console.info('onSuccessItem:responce', response,'status', status);
    };
    uploader.onErrorItem = function(fileItem, response, status, headers) {
        //console.info('onErrorItem', fileItem, response, status, headers);
    };
    uploader.onCancelItem = function(fileItem, response, status, headers) {
        //console.info('onCancelItem', fileItem, response, status, headers);
        console.log('item cancelled');
        updateHeight();
    };
    uploader.onCompleteItem = function(fileItem, response, status, headers) {

        // remove 'uploading' class of uploaded file
        var completeItemIndex = uploader.getIndexOfItem(fileItem);
        var completedItems = document.getElementsByClassName('file');
        completedItems[completeItemIndex].className = "file";

        // Checks if file already exists
        // if (!(status == '')) {
        //     var el = document.createElement( 'html' );
        //     el.innerHTML = response;
        //     var errorMessage = errorMessage = JSON.stringify(el.getElementsByTagName( 'pre' )[0].innerHTML.split('\n')[0]);            
        //     if (errorMessage == '"Error: fileExists"') {
        //         fileItem.fileExists = true;
        //     } else {
        //         fileItem.fileExists = false;
        //     }            
        // }
        if (status == 409){
            fileItem.fileExists = true;
        }

        // Shows error if session has expired. Right now if this happens the user
        // is not kicked and page stays open. If they upload something in this 
        // state the front end will happily comply but the back end will not. 
        // Results in a 'false' upload, as nothing is stored. 
        if (status !== 204 && !fileItem.isCancel) {
            fileItem.isTimeout = true; // isTimeout is my own property, not from docs
        }
    };


}]);

    // fixing scroll issues between sidebar and rest of page
    function updateHeight(){
        setTimeout(function () {
            
            var bodyHeight = $("#main-content").outerHeight(true);
            console.log("body height: " + bodyHeight);
            

            $("#side-bar").css("min-height", bodyHeight);
            $("#main-body").css("height", bodyHeight);
            console.log("main height: " +  bodyHeight);

        }, 100);
     };

     // expanding and collapsing of side bar tabs on desktop side
     var acc = document.getElementsByClassName("accordion");
     var i;
     for(i = 0; i < acc.length; i++){
        acc[i].onclick=function(){
            this.classList.toggle("active");
            this.nextElementSibling.classList.toggle("show");
        }
     }
