jQuery(document).ready(function($){

	// desktop menu
	$('#instructions-trigger').click(function(){
		$('header').attr('id', 'show-instructions');
	});

	$('#about-trigger').click(function(){
		$('header').attr('id', 'show-about');
	});


	$('#contact-trigger').click(function(){
		$('header').attr('id', 'show-contact');
	});

	$('#close-trigger').click(function(){
		$('header').attr('id', '');
	});

	$('#user-profile').click(function(){
		if( $('#account-options').css('visibility') == 'visible'){
			$('#account-options').css('visibility', 'hidden');
		}else{
			$('#account-options').css('visibility', 'visible');
		} 
	});

	// mobile menu

	$('#mobile-menu-trigger').click(function(){
		if( $('#mobile-topbar-expanded').hasClass('collapsed-container')){
			$('#mobile-topbar-expanded').removeClass('collapsed-container');
		}else{
			$('#mobile-topbar-expanded').addClass('collapsed-container');
		}
	});

	$('#mobile-account-trigger').click(function(){
		if( $('#mobile-account-container').hasClass('collapsed-container')) {
			$('#mobile-account-container').removeClass('collapsed-container');
		}else{
			$('#mobile-account-container').addClass('collapsed-container');
		}
	});

	$('#mobile-instruction-trigger').click(function(){
		if( $('#mobile-instruction-container').hasClass('collapsed-container')) {
			$('#mobile-instruction-container').removeClass('collapsed-container');
		}else{
			$('#mobile-instruction-container').addClass('collapsed-container');
		}
	});

	$('#mobile-about-trigger').click(function(){
		if( $('#mobile-about-container').hasClass('collapsed-container')) {
			$('#mobile-about-container').removeClass('collapsed-container');
		}else{
			$('#mobile-about-container').addClass('collapsed-container');
		}
	});

	$('#mobile-contact-trigger').click(function(){
		if( $('#mobile-contact-container').hasClass('collapsed-container')) {
			$('#mobile-contact-container').removeClass('collapsed-container');
		}else{
			$('#mobile-contact-container').addClass('collapsed-container');
		}
	});

});