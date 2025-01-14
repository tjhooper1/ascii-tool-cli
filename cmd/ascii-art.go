package cmd

var asciiArtOptions = []string{
  `⣿⣿⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⣿⣿⡿⠋⠟⠉⠉⢻⡿⢿⣿⣅⠈⠙⢿⣷⠀⠙⣿⣿⣿⠟⠉⢙⣾⣿⡿⠋⣹⣿⣷⣿⣿⢿⣿⣿⣿⣿⣿⡇⠀⣿⣿⠀⣿⣿⠀⡇⠁
⣿⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⡇⠘⣿⠀⣿⣿⠁⠀⠀⠀⠀⠈⠓⠦⠄⠉⠓⠀⠁⠹⡆⠀⠈⣿⠃⠀⣠⠿⠛⠉⢀⣴⣿⡿⠟⠛⠛⠛⠿⣿⣿⣿⣿⠀⢠⣿⣿⠁⣿⡏⠀⡇⠀
⣿⣿⣿⣿⣿⣿⣿⡆⣿⠻⢿⣿⣷⣶⣿⣦⣿⣿⣄⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢤⡀⠀⠀⠚⠁⠀⣠⣴⣿⡋⠀⠀⠀⣀⣤⣴⣶⣿⣿⣿⣿⠆⢸⣿⣿⣄⢿⡇⠀⡇⠀
⣿⣿⡿⠋⠟⣿⣿⡷⢿⡆⠀⠈⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣇⠀⠀⠀⠈⠉⠉⠉⠉⠁⠐⠞⠁⠀⠲⣿⣟⣻⣿⣧⣴⣿⣿⢿⡿⣿⣧⡠⡇⠠
⣿⠿⠇⠀⠀⣿⠛⢷⣜⣯⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⠄⠀⠀⣠⡤⠀⠀⠀⠀⣸⣆⡀⠀⠀⠀⢀⠀⠘⣯⠉⠀⠀⠀⠀⠀⠉⠙⢛⣟⢻⣉⣵⣿⣿⣿⣧⣶⡷⠺
⣿⠀⠄⠀⠐⣿⠄⠀⠉⠙⠛⠒⠒⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡞⠁⠀⢀⡾⠃⠀⠀⣠⣶⣋⣭⣤⠷⢦⡀⠀⠈⣇⠀⢸⣷⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠻⣿⣿⣿⣿⣯⣇⠰⠀
⣅⡀⣄⣤⣤⣿⢦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡟⠀⠀⢀⡾⠀⠀⢠⡞⠁⠀⠉⠉⠀⠀⠀⠹⡆⠀⢸⡀⠀⢻⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠛⠿⣿⡿⠿⠷⠶⠶
⣀⠈⠀⠀⢸⣏⠀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⠀⠀⡾⠀⠀⣠⢿⠃⠀⣰⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⠀⣸⣇⠀⠘⣇⢠⡀⠀⠀⠀⠀⢰⣄⠀⠀⠀⠀⠀⢿⡄⠀⠀⠀
⠹⠆⠀⠀⠀⢿⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣼⠃⢀⣼⠃⣠⠞⠁⡟⢀⡼⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⡴⢣⠙⣦⠀⢻⡄⢻⡆⠀⠀⠀⠀⠙⣷⠛⠛⠛⢻⣾⣿⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠉⠙⠦⣤⣤⠀⠀⠀⠀⢠⡶⣿⠇⣠⣾⡷⠋⢁⡴⣾⡷⠯⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⡿⠶⢿⣄⠈⠙⣮⣳⠀⢻⡄⠀⠀⠀⠀⠘⢳⣄⠀⢸⣿⣿⡿⠿⠿
⠀⠀⠀⠀⠀⠀⢀⣴⠛⠉⠀⠀⠀⢀⡴⠋⠀⣾⠄⢹⣿⠀⠀⠈⠀⠉⠀⠀⠘⠃⠀⠀⠀⠀⠀⠀⠀⠀⠰⠋⠀⠀⠸⠈⠀⠀⢹⣿⡀⠈⣧⠀⠀⠀⠀⠀⠀⠉⠙⣿⣯⠟⠋⠉⠉
⠀⠀⠀⠀⠀⠀⠙⠛⠓⠒⣶⣆⣴⠯⠤⠶⠚⢹⣆⣾⢹⠀⠀⠀⣴⣖⣒⢶⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⣖⡟⡓⢦⠀⠀⢸⡿⢷⡌⣿⠀⠀⠀⢀⣀⣀⣀⡽⠟⡇⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⡁⠀⠀⠀⠀⠀⠀⢸⣿⠿⣟⡆⠀⣼⢤⣼⣿⣷⢿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⢾⣿⣿⣇⠘⡆⠀⣿⣿⣨⣿⣟⠀⠀⢀⡋⠉⠻⣦⡀⠀⡏⠀⠱⣤⠀
⠀⠀⠀⣄⣀⡀⠀⠀⠈⠁⠈⠙⣷⠀⠀⠀⠀⢻⡇⢠⡘⣿⠀⣹⡄⠻⡿⠟⠘⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠿⣿⣣⡼⠁⢹⡏⠀⣿⡜⣿⠀⠀⠈⢻⡶⠶⠾⠟⠂⣇⣀⡸⠛⠁
⠀⠀⢀⣤⣬⡭⣗⣦⡀⠀⠀⢠⠇⣠⣾⠃⠀⠸⡇⠀⢱⣿⡼⠋⢳⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⡄⠀⣼⠃⢀⡿⠀⢻⣷⣤⢀⣼⠃⠀⠀⠀⠀⣿⠟⠀⠀⠀
⠀⠀⠈⠀⠀⠀⢹⡇⠙⣆⠀⠚⠋⠁⠹⣤⢤⡀⠹⣄⠀⣿⠁⠀⠀⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣷⠀⢻⣠⣾⡿⣤⠟⣸⣿⡏⠁⠀⠀⠀⢠⡿⠃⠀⠀⠀⠁
⠀⠀⠀⠀⠀⠀⢸⡇⠀⠸⡶⢤⣶⣶⣶⣷⣶⣷⣤⣈⣳⣧⠀⠀⠀⣿⠀⠀⠀⠀⠀⢀⣀⡠⠤⢤⣀⡀⠀⠀⠀⠀⢰⠇⠀⢸⣿⣿⡶⠷⢶⣾⣿⡇⠀⠀⠀⢰⣾⠁⠀⠠⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢸⡇⠀⠀⣿⡏⠉⠉⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⡿⠀⠀⠀⠐⠛⠁⠀⠀⠀⠀⠀⠈⠛⠂⠀⣠⠏⠀⣠⣿⣯⣄⣠⣄⣠⣿⢹⡇⠀⠀⠀⣹⠇⠀⡀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢸⡇⠀⠀⣿⠇⠀⣀⣽⡿⠿⠿⠿⠿⠛⠛⠛⣦⡐⣇⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⢋⣠⣾⠟⠋⠉⠉⠉⠛⠿⠿⠾⠧⣤⣠⡼⡇⠀⣠⠙⠀⠃⠀⠀
⠀⠀⠀⠀⠀⠀⠈⡇⠀⠀⣿⣀⡾⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠘⣷⠘⣿⣓⢶⢤⣤⣄⣀⣀⣠⣀⣤⣤⣾⣿⣿⣿⣿⣥⡄⠀⠐⠒⠶⣦⣄⠀⠀⠀⣹⡏⢸⠁⢰⠁⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⣿⡟⠀⣀⣤⠤⢤⣤⣀⠀⠀⠀⠀⠀⢹⡇⢻⡊⠉⠉⠉⠹⠿⠿⠿⠿⠿⠿⠿⠿⠿⠳⣿⡏⠀⠀⠀⠀⠀⠀⣹⡷⢾⣻⣿⣧⡄⠀⡎⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡇⠀⢀⡟⣴⡿⠟⠉⢩⣳⣶⣌⢷⡄⠀⠀⠀⣸⣿⣸⣧⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡿⠛⠒⣶⡦⢤⣄⡾⠋⣠⣿⠿⠒⢻⣇⠀⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⣿⠏⠀⠀⠀⠀⠀⠀⠙⠳⠿⣤⡴⠋⠁⠀⢀⣀⣈⣙⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣥⣖⠋⠁⠀⠀⠀⠹⣟⠁⠀⠀⠀⠀⢹⡞⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠇⠀⠀⠀⠀⠀⣀⣉⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣏⣿⠒⠋⠀⠀⠀⠀⠈⠀⠀⠀⠀⠀⡼⡇⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⣨⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣯⠒⠒⠀⠀⠀⠀⠀⠀⠀⠀⣠⡞⠁⢸⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⣇⠱⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠒⢋⡵⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡟⢯⡑⠲⠀⠀⠀⠀⣠⣤⣴⡾⠭⠶⠚⣿⡄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⢸⡆⠀⠉⠙⠒⠒⢶⡶⠚⠋⠉⠉⠛⠲⠦⣴⠞⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡧⠶⠉⣷⢦⣤⠶⢿⡁⠀⣿⠃⠀⠀⠀⠻⡇⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢱⣀⣀⠜⢷⡀⠀⣀⣀⣠⣼⣇⠀⠀⠀⠀⠀⠀⠀⣻⡄⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢀⣴⣟⣁⢀⣀⢀⣸⠇⠀⣿⠀⠀⠀⠀⠀⣧⠀⠀⠀⠀⠀⠀
⠁⠀⠀⠀⠀⠀⠀⢸⠐⠒⠀⠀⡏⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠛⣿⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⠛⠉⠉⠉⠉⠉⠉⠉⠉⢍⡉⠙⢿⣟⠛⠽⣿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠘⠀⠀⠀⠀⣿⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⣓⣶⣶⣶⣶⣒⣲⣖⠒⢒⣚⣉⣁⣀⣀⣀⣀⣀⣀⣀⡀⠀⠀⠉⣦⣄⡙⣷⣖⣾⡄⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⠏⠉⠉⠉⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⢸⣏⡀⠀⠀⣙⣏⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠙⢿⣧⣄⠀⠀⠀⠀
⠠⠄⠀⠐⠀⠒⠒⠚⠛⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢧⠃⠀⠀⠈⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠳⢄⠀⠉⠿⣍⠻⠓⠒
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡶⠶⠆⠖⠲⢻⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠤⠀⠈⠻⣦⠀`,
  `⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡾⠋⠉⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠃⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⣠⠖⠲⢤⡖⠒⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⡏⠀⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⡏⠀⠀⠀⠀⠀⢀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣀⠀⠀⢸⠀⠀⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⠸⣄⠀⠁⣠⠞⠉⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣠⣤⣤⣤⣤⠀⠀
⠀⡞⠉⠻⠁⢹⠀⠀⡏⠀⠀⠀⠀⢸⠃⠀⠀⠀⠀⠀⠀⠀⠀⠹⣶⠋⠀⠀⠀⠀⣀⡤⠴⠒⠊⠉⠉⠀⠀⣿⣿⣿⠿⠋⠀⠀
⠀⠳⢤⡀⠀⡞⠁⠀⡇⠀⠀⢀⡠⠼⠴⠒⠒⠒⠒⠦⠤⠤⣄⣀⠀⢀⣠⠴⠚⠉⠀⠀⠀⠀⠀⠀⠀⠀⣼⠿⠋⠁⠀⠀⠀⠀
⠀⠀⠀⠈⠷⡏⠀⠀⣇⠔⠂⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢨⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⣀⡤⠖⠋⠁⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢰⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⠤⠒⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢀⡟⠀⣠⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⢻⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣤⣤⡤⠤⢴
⠀⠀⠀⠀⠀⠀⣸⠁⣾⣿⣀⣽⡆⠀⠀⠀⠀⠀⠀⠀⢠⣾⠉⢿⣦⠀⠀⠀⢸⡀⠀⠀⢀⣠⠤⠔⠒⠋⠉⠉⠀⠀⠀⠀⢀⡞
⠀⠀⠀⠀⠀⢀⡏⠀⠹⠿⠿⠟⠁⠀⠰⠦⠀⠀⠀⠀⠸⣿⣿⣿⡿⠀⠀⠀⢘⡧⠖⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡼⠀
⠀⠀⠀⠀⠀⣼⠦⣄⠀⠀⢠⣀⣀⣴⠟⠶⣄⡀⠀⠀⡀⠀⠉⠁⠀⠀⠀⠀⢸⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠁⠀
⠀⠀⠀⠀⢰⡇⠀⠈⡇⠀⠀⠸⡾⠁⠀⠀⠀⠉⠉⡏⠀⠀⠀⣠⠖⠉⠓⢤⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⠃⠀⠀
⠀⠀⠀⠀⠀⢧⣀⡼⠃⠀⠀⠀⢧⠀⠀⠀⠀⠀⢸⠃⠀⠀⠀⣧⠀⠀⠀⣸⢹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡰⠃⠀⠀⠀
⠀⠀⠀⠀⠀⠈⢧⡀⠀⠀⠀⠀⠘⣆⠀⠀⠀⢠⠏⠀⠀⠀⠀⠈⠳⠤⠖⠃⡟⠀⠀⠀⢾⠛⠛⠛⠛⠛⠛⠛⠛⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠙⣆⠀⠀⠀⠀⠈⠦⣀⡴⠋⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⠙⢦⠀⠀⠘⡇⠀⠀⠀⠀⠀⠀⢀⣀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢠⡇⠙⠦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⠴⠋⠸⡇⠈⢳⡀⠀⢹⡀⠀⠀⠀⢀⡞⠁⠉⣇⣀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡼⣀⠀⠀⠈⠙⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀⠀⠀⠀⣷⠴⠚⠁⠀⣀⣷⠀⠀⠀⢠⠇⠀⠀⠀⠀⠀⣳
⠀⠀⠀⠀⠀⠀⡴⠁⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣆⡴⠚⠉⠉⠀⠀⠀⠀⢸⠃⣀⣠⠤⠤⠖⠋
⣼⢷⡆⠀⣠⡴⠧⣄⣇⠀⠀⠀⠀⡴⠚⠙⠲⠞⠛⠙⡆⠀⠀⠀⠀⠀⢀⡇⣠⣽⢦⣄⢀⣴⣶⠀⠋⠉⠀⠀⠀⠀⠀⠀⠀⠀
⡿⣼⣽⡞⠁⠀⠀⠀⢹⡀⠀⠀⠀⢹⠀⠀⠀⠀⠀⠀⣸⠀⠀⠀⠀⠀⣼⠉⠁⠀⠀⢠⢟⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣷⠉⠁⢳⠀⠀⠀⠀⠈⣧⠀⠀⠀⠀⠙⢦⠀⠀⠀⡠⠁⠀⠀⠀⠀⣰⠃⠀⠀⠀⠀⠏⠀⠀⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠹⡆⠀⠈⡇⠀⠀⠀⠀⠘⣆⠀⠀⠀⠀⠀⠹⣧⠞⠁⠀⠀⠀⠀⣰⠃⠀⠀⠀⠀⠀⠀⠀⣸⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢳⡀⠀⠙⠀⠀⠀⠀⠀⠘⣆⠀⠀⠀⠀⠀⡇⠀⠀⠀⠀⠀⣰⠃⠀⠀⠀⠀⢀⡄⠀⢠⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢳⡀⣰⣀⣀⣀⠀⠀⠀⠘⣦⣀⠀⠀⠀⡇⠀⠀⠀⢀⡴⠃⠀⠀⠀⠀⠀⢸⡇⢠⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠉⠉⠀⠀⠈⠉⠉⠉⠙⠻⠿⠾⠾⠻⠓⢦⠦⡶⡶⠿⠛⠛⠓⠒⠒⠚⠛⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,

`
 
`,
  // Add more ASCII art options here...
}

var asciiArtTitles = []string{
  "Studio Ghibli",
  "Heart",
  "Face",
  // Add titles for each ASCII art here...
}
