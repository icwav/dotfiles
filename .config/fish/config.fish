if status is-interactive
    # Commands to run in interactive sessions can go here
# sudo !! function
function sudo
    if test "$argv" = !!
        echo sudo $history[1]
        eval command sudo $history[1]
    else
        command sudo $argv
    end
end

set -g theme_powerline_fonts no
set -g theme_nerd_fonts yes
set -g theme_display_date no
set -g theme_color_scheme nord
set -g theme_title_display_user no
set -g theme_display_cmd_duration no
set -g theme_title_display_process no

end
