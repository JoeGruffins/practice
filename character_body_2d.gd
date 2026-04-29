extends CharacterBody2D                                                                                                                                                              
																																													   
const SPEED = 200.0
const JUMP_VELOCITY = -400.0                                                                                                                                                         
const GRAVITY = 980.0                                           

func _physics_process(delta):                                                                                                                                                        
		# Gravity - pulls you down when not on the floor
		if not is_on_floor():                                                                                                                                                          
				velocity.y += GRAVITY * delta                     
																																													   
		# Jump - only when on the floor                           
		if Input.is_action_just_pressed("ui_accept") and is_on_floor():
				velocity.y = JUMP_VELOCITY                                                                                                                                             
																																													   
		# Horizontal movement                                                                                                                                                          
		var direction = Input.get_axis("ui_left", "ui_right")                                                                                                                          
		velocity.x = direction * SPEED                            

		move_and_slide()
